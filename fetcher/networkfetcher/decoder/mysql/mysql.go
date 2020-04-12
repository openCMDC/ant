package mysql

import (
	"ant/core"
	"ant/fetcher/networkfetcher/base"
	"ant/fetcher/networkfetcher/decoder"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

const mysqlName = "mysql"

type mysql struct {
	stmtMap   map[uint32]*Stmt
	pktChan   chan *packet
	closeChan chan struct{}
	conn      *base.TCPConn

	lastC2SPkt *packet
	S2CPkts    []*packet
}

type packet struct {
	isClient2ServerFlow bool
	seq                 int
	length              int
	payload             []byte
	ts                  int64
}

var mysqlTable core.Table
var recordMap = make(map[byte]string)

func init() {
	decoder.RegisterDecoder(mysqlName, NewMysqlDecoder)
	fs := make([]*core.Field, 0)
	fs = append(fs, &core.Field{Name: "sourceIpAndPort", DataType: core.String})
	fs = append(fs, &core.Field{Name: "dstIpAndPort", DataType: core.String})
	fs = append(fs, &core.Field{Name: "type", DataType: core.String})
	fs = append(fs, &core.Field{Name: "sql", DataType: core.String})
	fs = append(fs, &core.Field{Name: "parameters", DataType: core.String})
	fs = append(fs, &core.Field{Name: "rt", DataType: core.Number})
	fs = append(fs, &core.Field{Name: "requestSize", DataType: core.Number})
	fs = append(fs, &core.Field{Name: "responseSize", DataType: core.Number})
	mysqlTable = core.Table{
		Name:   "mysql",
		Desc:   "mysql",
		Fields: fs,
	}

	recordMap[3] = "query"  //COM_QUERY 增删改查似乎都是走这两个接口
	recordMap[23] = "query" // COM_STMT_EXECUTE 增删改查似乎都是走这两个接口
}

func (m *mysql) Name() string {
	return mysqlName
}

func (m *mysql) Parse(conn *base.TCPConn, senBackChan chan<- *core.Row) {
	c2s := conn.C2SStream()
	s2c := conn.S2CStream()

	//time.Sleep(time.Millisecond * 50)
	m.conn = conn
	go getPktCyclically(decoder.NewNamedReader(mysqlName, c2s), true, m.pktChan, m.closeChan)
	go getPktCyclically(decoder.NewNamedReader(mysqlName, s2c), false, m.pktChan, nil)

	m.startProcess(senBackChan)
}

func (m *mysql) startProcess(senBackChan chan<- *core.Row) {
	goOn := true
	for goOn {
		select {
		case pkt := <-m.pktChan:
			if pkt == nil {
				continue
			}
			if pkt.length != 0 {
				if pkt.isClient2ServerFlow {
					m.resolveClientPacket(pkt, senBackChan)
				} else {
					m.resolveServerPacket(pkt)
				}
			}
		case <-m.closeChan:
			goOn = false
		}
	}
	close(m.pktChan)
	close(m.closeChan)
}

func (m *mysql) resolveClientPacket(pkt *packet, senBackChan chan<- *core.Row) {

	payload, seq := pkt.payload, pkt.seq

	if m.lastC2SPkt != nil {
		p := m.lastC2SPkt.payload
		row := m.buildBaseRow()
		index := p[0]
		t := recordMap[index]
		if len(t) > 0 {
			row.Content = append(row.Content, t)
			//获取sql语句
			if index == COM_STMT_EXECUTE {
				stmtID := binary.LittleEndian.Uint32(p[1:5])
				sql := m.stmtMap[stmtID]
				row.Content = append(row.Content, sql)
				//todo 解析参数
				row.Content = append(row.Content, "")
			} else {
				sql := string(p[1:])
				row.Content = append(row.Content, sql)
				row.Content = append(row.Content, "")
			}
			//record response time
			if len(m.S2CPkts) == 0 {
				row.Content = append(row.Content, 0)
			} else {
				row.Content = append(row.Content, m.S2CPkts[0].ts-m.lastC2SPkt.ts)
			}
			//record request size
			row.Content = append(row.Content, len(p))
			//record response size
			repSize := 0
			for _, pk := range m.S2CPkts {
				repSize += len(pk.payload)
			}
			row.Content = append(row.Content, repSize)
			senBackChan <- row
		}
	}

	//clear
	m.lastC2SPkt = pkt
	m.S2CPkts = m.S2CPkts[0:0]

	var msg string
	switch payload[0] {

	case COM_INIT_DB:
		msg = fmt.Sprintf("USE %s;\n", payload[1:])
	case COM_DROP_DB:

		msg = fmt.Sprintf("Drop DB %s;\n", payload[1:])
	case COM_CREATE_DB, COM_QUERY:

		statement := string(payload[1:])
		msg = fmt.Sprintf("%s %s", ComQueryRequestPacket, statement)
	case COM_STMT_PREPARE:

		for {
			select {
			case packet, ok := <-m.pktChan:
				if !ok {
					log.Println("ERR : Not found stm packet")
					return
				}
				if packet.seq == seq+1 {
					//fetch stm id
					stmtID := binary.LittleEndian.Uint32(packet.payload[1:5])
					stmt := &Stmt{
						ID:    stmtID,
						Query: string(payload[1:]),
					}

					//record stm sql
					m.stmtMap[stmtID] = stmt
					stmt.FieldCount = binary.LittleEndian.Uint16(packet.payload[5:7])
					stmt.ParamCount = binary.LittleEndian.Uint16(packet.payload[7:9])
					stmt.Args = make([]interface{}, stmt.ParamCount)

					msg = PreparePacket + stmt.Query
				}
			case <-time.After(5 * time.Second):
				log.Println("ERR : Not found stm packet")
				return
			}
		}
	case COM_STMT_SEND_LONG_DATA:

		stmtID := binary.LittleEndian.Uint32(payload[1:5])
		paramId := binary.LittleEndian.Uint16(payload[5:7])
		stmt, _ := m.stmtMap[stmtID]

		if stmt.Args[paramId] == nil {
			stmt.Args[paramId] = payload[7:]
		} else {
			if b, ok := stmt.Args[paramId].([]byte); ok {
				b = append(b, payload[7:]...)
				stmt.Args[paramId] = b
			}
		}
		return
	case COM_STMT_RESET:

		stmtID := binary.LittleEndian.Uint32(payload[1:5])
		stmt, _ := m.stmtMap[stmtID]
		stmt.Args = make([]interface{}, stmt.ParamCount)
		return
	case COM_STMT_EXECUTE:

		var pos = 1
		stmtID := binary.LittleEndian.Uint32(payload[pos : pos+4])
		pos += 4
		var stmt *Stmt
		var ok bool
		if stmt, ok = m.stmtMap[stmtID]; ok == false {
			log.Println("ERR : Not found stm id", stmtID)
			return
		}

		//params
		pos += 5
		if stmt.ParamCount > 0 {

			//（Null-Bitmap，len = (paramsCount + 7) / 8 byte）
			step := int((stmt.ParamCount + 7) / 8)
			nullBitmap := payload[pos : pos+step]
			pos += step

			//Parameter separator
			flag := payload[pos]

			pos++

			var pTypes []byte
			var pValues []byte

			//if flag == 1
			//n （len = paramsCount * 2 byte）
			if flag == 1 {
				pTypes = payload[pos : pos+int(stmt.ParamCount)*2]
				pos += int(stmt.ParamCount) * 2
				pValues = payload[pos:]
			}

			//bind params
			err := stmt.BindArgs(nullBitmap, pTypes, pValues)
			if err != nil {
				log.Println("ERR : Could not bind params", err)
			}
		}
		msg = string(stmt.WriteToText())
	default:
		return
	}
	fmt.Println(msg)
}

func (m *mysql) resolveServerPacket(pkt *packet) {

	m.S2CPkts = append(m.S2CPkts, pkt)
	payload, _ := pkt.payload, pkt.seq
	var msg = ""
	if len(payload) == 0 {
		return
	}
	switch payload[0] {

	case 0xff:
		//errorCode  := int(binary.LittleEndian.Uint16(payload[1:3]))
		//errorMsg,_ := ReadStringFromByte(payload[4:])
		//
		//msg = GetNowStr(false)+"%s Err code:%s,Err msg:%s"
		//msg  = fmt.Sprintf(msg, ErrorPacket, strconv.Itoa(errorCode), strings.TrimSpace(errorMsg))

	case 0x00:
		//var pos = 1
		//l,_ := LengthBinary(payload[pos:])
		//affectedRows := int(l)
		//
		//msg += GetNowStr(false)+"%s Effect Row:%s"
		//msg = fmt.Sprintf(msg, OkPacket, strconv.Itoa(affectedRows))

	default:
		return
	}

	fmt.Println(msg)
}

func (m *mysql) buildBaseRow() *core.Row {
	res := new(core.Row)
	res.Meta = &mysqlTable
	res.Content = make([]interface{}, 0)
	res.Content = append(res.Content, m.conn.GetClientAddr().String())
	res.Content = append(res.Content, m.conn.GetServerAddr().String())
	return res
}

func NewMysqlDecoder() decoder.Interface {
	return &mysql{
		stmtMap:   make(map[uint32]*Stmt),
		pktChan:   make(chan *packet),
		closeChan: make(chan struct{}),
		S2CPkts:   make([]*packet, 0),
	}
}
