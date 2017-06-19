package enviroment

import (
	"database/sql"
	"fmt"

	seelog "github.com/cihub/seelog"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

type ConfigFile struct {
	Database string
	Port     string
	LogDes   string
}

type Config struct {
	Db     *sql.DB
	Port   string
	LogDes string
	Err    error
}

var Logger seelog.LoggerInterface

func loadAppConfig(appConfig string) {

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

//EnvBuild需要正确的解析文件并且初始化DB和Redis的连接。。
func EnvBuild(filepath string) Config {

	var tmp ConfigFile
	var conf Config

	if _, err := toml.DecodeFile(filepath, &tmp); err != nil {
		conf.Err = err
		return conf
	}

	conf.Port = tmp.Port
	conf.LogDes = tmp.LogDes

	//open db
	db, err := sql.Open("mysql", tmp.Database)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.Ping()
	if err != nil {
		conf.Err = err
		return conf
	}
	conf.Db = db

	DisableLog()
	loadAppConfig(conf.LogDes)

	conf.Err = nil
	return conf
}
