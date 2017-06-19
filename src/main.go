package main

import (
	"fmt"
	"os"
	. "wbproject/httpserver/src/envbuild"
	. "wbproject/httpserver/src/process"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0PR1"
var filepath = "../etc/config.toml"
var cfg Config

func main() {

	args := os.Args

	//版本号
	if len(args) == 2 && (args[1] == "-v") {
		fmt.Println("现在的版本是【", version, "】")
		return
	}

	cfg = EnvBuild(filepath)
	if cfg.Err != nil {
		panic(cfg.Err)
	}

	router := gin.Default()
	router.POST("/register-student", RegStudentHandler)
	router.POST("/register-class", RegClassHandler)
	router.GET("/get-class-total-score/:sid", GetScoreHandler)
	router.GET("/get-class-total-score", GetTeacherHandler)
	router.Run(cfg.Port)
}

func RegStudentHandler(c *gin.Context) {

	RegStudent(&cfg, c)
}

func RegClassHandler(c *gin.Context) {

	RegClass(&cfg, c)
}

func GetScoreHandler(c *gin.Context) {

	GetScore(&cfg, c)
}

func GetTeacherHandler(c *gin.Context) {
	GetTeacher(&cfg, c)
}
