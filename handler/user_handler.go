package handler

import (
	"log"
	"net/http"
	"restful/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取全部数据处理事件
func GetAll(context *gin.Context) {
	user := model.UserModel{}

	users := user.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}


// 根据id获取单条数据
func GetOne(context *gin.Context) {
	id := context.Param("id")

	i, e := strconv.Atoi(id)
	if e!=nil {
		log.Panicln("id不是int类型", e.Error())
	}

	user := model.UserModel{
		Id: i,
	}

	u := user.FindById()
	context.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

// 插入数据
func Insert(context *gin.Context){
	user := model.UserModel{}

	var id int = -1
	if e := context.ShouldBindJSON(&user); e == nil {
		id = user.Insert()
	}

	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// 删除数据
func DeleteOne(context *gin.Context) {
	id := context.Param("id")

	i, e := strconv.Atoi(id)
	if e!=nil {
		log.Panicln("id不是int类型", e.Error())
	}

	user := model.UserModel{
		Id: i,
	}

	res := user.DeleteOne()
	context.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
