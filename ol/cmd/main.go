/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 21:39:11
 * @FilePath: \OpenLibrary\ol\cmd\main.go
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"ol/api/httpserver"
	httpSvrImpl "ol/api/httpserver/impl"
	"ol/config"
	cfgImpl "ol/config/impl"
	datasqlite "ol/internal/datasqlite"
	sqliteData "ol/internal/datasqlite/data"
	actionImpl "ol/internal/datasqlite/impl"
	"ol/internal/dbgorm"
	"ol/internal/server"
	"ol/internal/server/impl"
	"ol/public/logger"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	// -d=true
	debug   = flag.Bool("d", false, "if true, program will print detail logs")
	workDir = flag.String("wd", "", "specified the program work directory")
	cfgFile = flag.String("c", "config.yaml", "specified the config file path")

	log = logrus.WithFields(logrus.Fields{
		"Main": "",
	})
)

func flagParse() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *debug {
		log.Logger.SetLevel(logrus.DebugLevel)
	}

	appFileDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if *workDir != "" {
		*workDir = filepath.Clean(*workDir)
		if !filepath.IsAbs(*workDir) {
			*workDir = filepath.Join(appFileDir, *workDir)
		}
	} else {
		*workDir = appFileDir
	}

	err := os.Chdir(*workDir)
	if err != nil {
		fmt.Printf("Error: can not change work directory to %s. %v.\n", *workDir, err)
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: can not reach the work directory. %v.\n", err)
		os.Exit(1)
	}

	if !filepath.IsAbs(*cfgFile) {
		newCfg := filepath.Join(wd, *cfgFile)
		cfgFile = &newCfg
	}
}

func main() {
	flagParse()

	defer func() {
		if err := recover(); err != nil {
			log.Errorln("main recover error:", err)
		}
	}()

	// load server config from file config.yaml
	cfg, err := config.LoadCentralServerConfig(*cfgFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: please config the config file first.\n")
			os.Exit(1)
		}
		fmt.Printf("Error: server load config failed. %v.\n", err)
		os.Exit(1)
	}

	// 开启日志
	logger.DefaultLogger()
	logrus.SetLevel(logrus.InfoLevel)

	log.Info("run at path: ", *workDir)

	// mysql 数据库
	var gormDB *gorm.DB
	// sqlite
	var sqliteHandler datasqlite.DataSqlite
	if cfg.Database.HasUsed() {
		gormDB, err = dbgorm.InitGormDB("mysql", cfg.Database)
		if err != nil {
			log.Error("InitGormDB error:", err)
			os.Exit(1)
		}
	} else {
		sqliteDb := actionImpl.NewAction()
		err = sqliteDb.CreateTable(sqliteData.Book)
		if err != nil {
			log.Error("CreateTable error:", err)
			os.Exit(1)
		}

		sqliteHandler = sqliteDb
	}

	svr := impl.NewServer(gormDB, sqliteHandler)
	err = svr.Start()
	if err != nil {
		log.Errorln("NewServer Start error:", err)
		return
	}

	// http server
	err = startHTTPServer(svr, cfg)
	if err != nil {
		log.Errorln("startHTTPServer error:", err)
		return
	}

	log.Info("service start success")

	waitForSignal()
	log.Info("server quit success.")
}

func startHTTPServer(server server.EventListener, cfg *cfgImpl.ServerConfig) error {
	var httpServer httpserver.GinServer
	httpServer = httpSvrImpl.NewGinHTTPServer(cfg.HTTPCfg.HTTPHost, cfg.HTTPCfg.HTTPPort)

	err := httpServer.SetEventListenServer(server)
	if err != nil {
		log.Errorln("SetEventListenServer error:", err)
		return err
	}

	err = httpServer.Start()
	if err != nil {
		log.Errorln("httpServer Start error:", err)
		return err
	}
	return nil
}

func waitForSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sig
}
