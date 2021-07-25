package main

import (
	"log"
	"net/http"
	"playground/blog-service/global"
	"playground/blog-service/internal/model"
	"playground/blog-service/internal/routes"
	"playground/blog-service/pkg/logger"
	"playground/blog-service/pkg/setting"
	"playground/blog-service/pkg/tracer"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setuptSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 部落格系統
// @version 1.0
// @description Go 語言程式設計之旅: 一起用 Go 做專案
// @termsOfService https://github.com/yakushou730/go-playground
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routes.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"localhost:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
