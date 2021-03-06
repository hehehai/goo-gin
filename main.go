package main

import (
	"fmt"
	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"log"
	"net/http"
)
// @title gin-example
// @verson 0.1.0
// @contact.name guanwei
// @contact.email riverhohai@gmail.com

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ReadTimeout
	writTimeout := setting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
