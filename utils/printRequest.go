package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const printerMacAddress = "08:13:F4:C4:34:53"
const fixedTagText = "DB23LPTP3"
const fixedQRLinkFormat = "https://example.org/%s"

func PrintRequest(c *gin.Context) {

	err := PrintTag(fixedTagText, fixedQRLinkFormat)
	if err != nil {
		SendJSONResponse(c, gin.H{"error": "Error printing"}, http.StatusInternalServerError)
		return
	}

	SendJSONResponse(c, gin.H{"message": "Print request sent successfully"}, http.StatusOK)
}
