package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../bean"
	"../handlers"
)

type option struct {
	err error
	data []bean.User
}

func getRes(option *option) (res bean.Res) {
	if option.err != nil {
		res.Code = 500
		res.Message = "error"
	} else {
		res.Code = 200
		res.Message = "Success"
		res.Data = option.data
	}
	return 
}

func Insert(c *gin.Context) {
	name := c.PostForm("name")
	age := c.DefaultPostForm("age", "17")
	phone := c.PostForm("phone")
	if name == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{"res": 1002, "message": "参数缺失", "ip": c.ClientIP()})
		return
	}
	var data bean.User
	data.Name = name
	data.Age = age
	data.Phone = phone
	err := handlers.Insert(data)
	c.JSON(http.StatusOK, getRes(&option{err: err}))
	
}

func Find(c *gin.Context) {
	name := c.PostForm("name")
	user, err := handlers.Find(name)
	c.JSON(http.StatusOK, getRes(&option{err: err, data: user}))
}

func Update(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	age := c.PostForm("age")
	if name == "" || phone == "" || age == "" {
		c.JSON(http.StatusOK, gin.H{"res": 1002, "message": "参数缺失", "ip": c.ClientIP()})
		return
	}
	err := handlers.Update(name, age, phone)
	c.JSON(http.StatusOK, getRes(&option{err: err}))
}

func Delete(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"res": 1002, "message": "参数缺失", "ip": c.ClientIP()})
		return
	}
	err := handlers.Delete(name)
	c.JSON(http.StatusOK, getRes(&option{err: err}))
}