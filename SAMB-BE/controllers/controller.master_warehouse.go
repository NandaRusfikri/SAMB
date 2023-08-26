package controllers

import (
	services "SAMB-BE/services/master_warehouse"
	"SAMB-BE/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerMasterWareHouse struct {
	Service services.ServiceMasterWareHouse
}

func InitControllerMasterWareHouse(g *gin.Engine, service services.ServiceMasterWareHouse) {
	handler := &ControllerMasterWareHouse{
		Service: service,
	}

	routeAPI := g.Group("/api/v1/")
	routeAPI.GET("/warehouses", handler.ListWareHouseController)
}

// ListWareHouseController godoc
// @Tags MasterWareHouse
// @Summary ListWareHouseController
// @Description ListWareHouseController
// @ID ListWareHouseController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/warehouses [get]
func (h *ControllerMasterWareHouse) ListWareHouseController(c *gin.Context) {

	res, err := h.Service.ListMasterWareHouseService()
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}

}
