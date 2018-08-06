package utils

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/system"
	"github.com/MEIGUOSHU/meiguoshu_api_server/tools/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/cihub/seelog"
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
