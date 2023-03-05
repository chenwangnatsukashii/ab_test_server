package response

import (
	"github.com/gin-gonic/gin"
	"line_china/common/constant"
	"net/http"
)

// ResultVO 统一http返回结果
type ResultVO struct {
	Msg     constant.ResponseMsg `json:"msg"`
	Success bool                 `json:"success"`
	Data    interface{}          `json:"data"`
}

// Success 成功返回
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ResultVO{Msg: constant.SuccessMsg, Success: true, Data: data})
}

// Failure 失败返回
func Failure(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusInternalServerError, ResultVO{Msg: constant.ErrorMsg, Success: false, Data: data})
}
