package utils

import (
	"github.com/MEIGUOSHU/web-api/packages/system"
	"github.com/MEIGUOSHU/web-api/tools/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/seelog"
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

func Str2Int(string2 string) (int) {
	i, e := strconv.Atoi(string2)
	if e != nil {
		seelog.Infof("string转int错误，%s", e)
		seelog.Flush()
	}
	return i
}
