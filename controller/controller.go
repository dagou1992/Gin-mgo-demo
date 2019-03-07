package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../bean"
	"../handlers"
)

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
	var res bean.Res
	if err != nil {
		res.Code = 500
		res.Message = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = 200
	res.Message = "Success"
	c.JSON(http.StatusOK, res)
	
}

func Find(c *gin.Context) {
	name := c.PostForm("name")
	var res bean.Res
	user, err := handlers.Find(name)
	if err != nil {
		res.Code = 500
		res.Message = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = 200
	res.Message = "Success"
	res.Data = user
	c.JSON(http.StatusOK, res)
}

func Update(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	age := c.PostForm("age")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"res": 1002, "message": "参数缺失", "ip": c.ClientIP()})
		return
	}
	var res bean.Res
	err := handlers.Update(name, age, phone)
	if err != nil {
		res.Code = 500
		res.Message = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = 200
	res.Message = "Success"
	c.JSON(http.StatusOK, res)
}

func Delete(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"res": 1002, "message": "参数缺失", "ip": c.ClientIP()})
		return
	}
	err := handlers.Delete(name)
	var res bean.Res
	if err != nil {
		res.Code = 500
		res.Message = "error"
		c.JSON(http.StatusOK, res)
		return
	}
	res.Code = 200
	res.Message = "Success"
	c.JSON(http.StatusOK, res)
}