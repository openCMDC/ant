package main

import (
	"ant/core"
	"ant/db"
	"ant/fetcher"
	_ "ant/fetcher/networkfetcher/decoder/http"
	_ "ant/fetcher/networkfetcher/decoder/mysql"
	_ "ant/fetcher/networkfetcher/decoder/redis"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func init() {
	// 设置日志格式为json格式
	//log.SetFormatter(&log.JSONFormatter{})
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	log.SetLevel(log.ErrorLevel)
}

func main() {
	antCtx := core.NewDefaultAntCtx("12132")

	store := db.NewTestStorage(antCtx)
	err := store.StartConsume()
	if err != nil {
		log.WithError(err).Errorf("start storage failed")
		return
	}

	fetcherManager := fetcher.NewFetcherManager(antCtx, store)
	//time.Sleep(time.Second * 20)
	err = fetcherManager.Start()
	if err != nil {
		log.WithError(err).Errorf("start network fetcher manager failed")
		return
	}
	http.ListenAndServe("localhost:6060", nil)

	//terminalChan := make(chan os.Signal)
	//signal.Notify(terminalChan, os.Interrupt, os.Kill)
	//s := <-terminalChan
	//fmt.Printf("terminate process, Got signal: %s \n", s)
}
