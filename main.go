package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
	"strconv"
	"./router"
	"./bean"
)

func main() {
	router := router.Init()

	config := bean.Config{}
	err := gcfg.ReadFileInto(&config, "./config/config.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
	}

	if config.Section.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	
	port := config.Section.Port

	router.Run(":" + strconv.Itoa(port))
}