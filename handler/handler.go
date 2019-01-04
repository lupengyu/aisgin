package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lupengyu/aisgin/dal/mysql"
	"github.com/lupengyu/aisgin/helper"
	"net/http"
)

func GetShipID(c *gin.Context) {
	shipIDs := mysql.GetShipID()
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipIDs)
}

func GetPosition(c *gin.Context) {
	id := c.Params.ByName("id")
	positions := mysql.GetPosition(id)
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, positions)
}

func GetShipInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	shipInfos := mysql.GetShipInfo(id)
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipInfos)
}