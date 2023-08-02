package main

import (
	"github.com/datumbrain/label-printer/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/print", utils.PrintRequest)
	r.Run(":7040")
}
