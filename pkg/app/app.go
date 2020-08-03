package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"walle-go/logger"
)

const (
	SUCCESS = 0
	ERROR   = -1
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUid(c *gin.Context) int {
	v, _ := c.Get("uid")
	if v != nil {
		uid := v.(int)
		return uid
	}
	return 0
}

func (g *Gin) ResponseError(msg string) {
	g.ResponseWithMsgAndCode(http.StatusOK, ERROR, msg, nil)
}

func (g *Gin) ResponseSuccess(data interface{}) {
	g.ResponseWithMsgAndCode(http.StatusOK, SUCCESS, "", data)
}

func (g *Gin) ResponseList(list interface{}, count int, enableCreate bool) {
	data := map[string]interface{}{
		"list":          list,
		"count":         count,
		"enable_create": enableCreate,
	}
	g.ResponseWithMsgAndCode(http.StatusOK, SUCCESS, "", data)
}

func (g *Gin) ResponseWithMsgAndCode(httpCode, errCode int, msg string, data interface{}) {
	var response = Response{
		Code:    errCode,
		Message: msg,
		Data:    data,
	}
	var resStr, _ = json.Marshal(response)
	var reqBody, _ = g.C.Get("http_request_raw_data")
	var reqHeader, _ = json.Marshal(g.C.Request.Header)
	logger.Infof("reqURL[%v] reqHeader[%v] reqRawData[%v] resBody[%v]", g.C.Request.RequestURI, string(reqHeader), reqBody, string(resStr))
	g.C.JSON(httpCode, response)
	return
}
