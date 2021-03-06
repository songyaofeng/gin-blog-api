package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/response"
	"go-gin-blog-api/models"
	"go-gin-blog-api/util"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	translateId := "ok"
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				translateId = "Error auth token."
			} else {
				data["token"] = token
				translateId = "ok"
			}

		} else {
			translateId = "Token error."
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	if translateId != "ok" {
		c.JSON(http.StatusOK, response.Failed{response.Fail, response.Translate(22)})
		return
	}

	response.Success(c)
}
