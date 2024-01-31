package shared

import "github.com/gin-gonic/gin"

type DefaultResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *gin.Context, status int, data interface{}, message *string) {
	ctx.JSON(status, DefaultResponse{
		Message: message,
		Data:    data,
	})
}
