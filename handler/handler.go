package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redisdrive "github.com/gomodule/redigo/redis"
	"github.com/lupengyu/aisgin/dal/mysql"
	"github.com/lupengyu/aisgin/dal/redis"
	"github.com/lupengyu/aisgin/helper"

	"encoding/json"
	"net/http"
)

func GetShipID(c *gin.Context) {
	shipIDs := mysql.GetShipID()
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipIDs)
}

func GetPosition(c *gin.Context) {
	id := c.Params.ByName("id")
	if redis.Client != nil {
		result, err := redisdrive.String(redis.Client.Do("GET", "position_"+id))
		if err == nil {
			helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, fmt.Sprint(result))
			return
		}
		positions := mysql.GetPosition(id)
		positionsByte, err := json.Marshal(positions)
		if err == nil {
			redisdrive.String(redis.Client.Do("SET", "position_"+id, positionsByte))
		}
		helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, positions)
		return
	}
	positions := mysql.GetPosition(id)
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, positions)
}

func GetShipInfo(c *gin.Context) {
	id := c.Params.ByName("id")
	if redis.Client != nil {
		result, err := redisdrive.String(redis.Client.Do("GET", "info_"+id))
		if err == nil {
			helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, fmt.Sprint(result))
			return
		}
		shipInfos := mysql.GetShipInfo(id)
		shipInfosByte, err := json.Marshal(shipInfos)
		if err == nil {
			redisdrive.String(redis.Client.Do("SET", "info_"+id, shipInfosByte))
		}
		helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipInfos)
		return
	}
	shipInfos := mysql.GetShipInfo(id)
	helper.BizResponse(c, http.StatusOK, helper.CodeSuccess, shipInfos)
}
