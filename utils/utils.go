package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/MEIGUOSHU/web-api/packages/system"
	"github.com/MEIGUOSHU/web-api/tools/msg"
)

//GenResponse genrate reponse ,json format
func Res(c *gin.Context, code int, data interface{}, message string) {
	if system.Config.Server.Mode == "release" {
		message = msg.GetResponseMessage(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}