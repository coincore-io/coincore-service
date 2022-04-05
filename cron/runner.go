package cron


import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

var RealTimeExecution  = time.Second * 10


func Run() {
	if beego.BConfig.RunMode == "dev" {
		RealTimeExecution = time.Second * 10
	}
	go func() {
		for {
			select {
			case <-time.Tick(RealTimeExecution):
				err := SyncGasPrice()
				if err != nil {
					logs.Error("run SyncGasPrice error %v", err)
				} else {
					logs.Info("run SyncGasPrice success")
				}
				err = RealMarketAssetPrice()
				if err != nil {
					logs.Error("run RealMarketAssetPrice error %v", err)
				} else {
					logs.Info("run RealMarketAssetPrice success")
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case now := <-time.Tick(time.Minute):
				if now.Hour() == 0 && now.Minute() == 1 {
					fmt.Println("定点执行任务")
				}
			}
		}
	}()
}

