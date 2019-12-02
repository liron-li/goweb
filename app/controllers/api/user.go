package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goweb/pkg/e"
	"goweb/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Profile(c *gin.Context) {
	token := c.Query("token")
	claims, err := util.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(401)
	}

	c.JSON(http.StatusOK, util.RetJson(e.Success, claims))
}

func Login(c *gin.Context) {

	data := make(map[string]interface{})
	code := e.InvalidParams

	c.JSON(http.StatusOK, util.RetJson(code, data))
}
