package helper

import (
	"github.com/gin-gonic/gin"
)

const (
	ReportHttpCode = "REPORT_HTTP_CODE"
	ReportBizCode  = "REPORT_BIZ_CODE"
)

const (
	CodeSuccess = int32(0)
	CodeFailed  = int32(1)
)

type GinResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func UpdateReportStatus(c *gin.Context, httpCode int, bizCode int32) {
	c.Set(ReportHttpCode, httpCode)
	c.Set(ReportBizCode, bizCode)
}

func NewGinResponse(code int32, result interface{}) *GinResponse {
	return &GinResponse{
		Code:    code,
		Message: "success",
		Data:    result,
	}
}

func BizResponse(c *gin.Context, httpCode int, bizCode int32, data interface{}) {
	UpdateReportStatus(c, httpCode, bizCode)

	resp := NewGinResponse(bizCode, data)
	c.JSON(httpCode, resp)
}

func AbortWithBizResponse(c *gin.Context, httpCode int, bizCode int32, data interface{}) {
	c.Abort()
	BizResponse(c, httpCode, bizCode, data)
}

func BizStatus(c *gin.Context, httpCode int) {
	bizCode := CodeSuccess
	if httpCode >= 500 {
		bizCode = CodeFailed
	}
	UpdateReportStatus(c, httpCode, bizCode)

	c.Status(httpCode)
}
