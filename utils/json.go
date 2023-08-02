package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendJSONResponse(ctx *gin.Context, data interface{}, statusCode int) {
	ctx.Header("Content-Type", "application/json")
	ctx.Writer.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("JSON Encoding Failed")
	}

	ctx.Writer.Write(jsonData)
}
