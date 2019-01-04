package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lupengyu/aisgin/dal/mysql"
	"github.com/lupengyu/aisgin/helper"
	"net/http"
)

func GetShipID(c *gin.Context) {
	//ctx := helper.GenCtxFromGin(c)
	shipIDs := mysql.GetShipID()
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipIDs)
}
