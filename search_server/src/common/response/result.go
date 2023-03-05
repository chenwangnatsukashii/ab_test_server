package response

import (
	"github.com/gin-gonic/gin"
	"line_china/common/constant"
	"net/http"
)

type ResultVO struct {
	Msg     constant.ResponseMsg `json:"msg"`
	Success bool                 `json:"success"`
	Data    interface{}          `json:"data"`
}

func Success(ctx *gin.Context, msg constant.ResponseMsg, data interface{}) {
	resp := &ResultVO{Msg: msg, Success: true, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func Failure(ctx *gin.Context, msg constant.ResponseMsg, data interface{}) {
	resp := &ResultVO{Msg: msg, Success: false, Data: data}
	ctx.JSON(http.StatusInternalServerError, resp)
}
