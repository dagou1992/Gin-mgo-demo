package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/gcfg.v1"
	"../bean"
)

func Connect() (mongo *mgo.Database) {
	config := bean.Config{}
	err := gcfg.ReadFileInto(&config, "./config/config.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
		panic(err.Error())
	}

	uri := string(config.MongoDB.Ip) + ":" + string(config.MongoDB.Port)

	session, err := mgo.Dial(uri)
	fmt.Println("Connected to", uri)
	if err != nil {
		fmt.Printf("Mongo connect failed")
		panic(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	mongo = session.DB(config.MongoDB.Database)
	return 
}