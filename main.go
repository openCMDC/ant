package main

import (
	"ant/fetcher"
	_ "ant/fetcher/networkfetcher/decoder/http"
	_ "ant/fetcher/networkfetcher/decoder/mysql"
	_ "ant/fetcher/networkfetcher/decoder/redis"
	_ "ant/fetcher/networkfetcher/decoder2/http"
	remote2 "ant/remote"
	"ant/task"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
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
	log.SetLevel(log.WarnLevel)
}

func main() {
	//_ := core.NewDefaultAntCtx()

	//store := db.NewTestStorage(antCtx)
	//err := store.StartConsume()
	//if err != nil {
	//	log.WithError(err).Errorf("start storage failed")
	//	return
	//}
	remote, err := remote2.NewFileRemote(`D:\openCMDC\ant\msgs`, 10*time.Second)
	if err != nil {
		log.WithField("error", err.Error()).Error("init remote failed")
		os.Exit(1)
	}
	taskManager := task.NewTaskManager(remote)
	fetcherManager, err := fetcher.NewFetcherManager(remote)
	if err != nil {
		log.WithField("error", err.Error()).Error("init fetcherManager failed")
		os.Exit(1)
	}
	fetcherManager.RegisterFechedRowProcessor(taskManager)
	err = fetcherManager.Start()
	if err != nil {
		log.WithField("error", err.Error()).Error("init fetcherManager failed")
		os.Exit(1)
	}
	//time.Sleep(time.Second * 20)
	//err := fetcherManager.Start()
	//if err != nil {
	//	log.WithError(err).Errorf("start network fetcher manager failed")
	//	return
	//}

	http.ListenAndServe("localhost:6060", nil)

	//terminalChan := make(chan os.Signal)
	//signal.Notify(terminalChan, os.Interrupt, os.Kill)
	//s := <-terminalChan
	//fmt.Printf("terminate process, Got signal: %s \n", s)
}
