package shared

import "github.com/gin-gonic/gin"

type DefaultResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponseWithMessage(ctx *gin.Context, status int, data interface{}, message string) {
	messagePointer := message
	ctx.JSON(status, DefaultResponse{
		Message: &messagePointer,
		Data:    data,
	})
}

func SendResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, DefaultResponse{
		Message: nil,
		Data:    data,
	})
}
