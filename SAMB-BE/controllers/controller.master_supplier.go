package controllers

import (
	services "SAMB-BE/services/master_supplier"
	"SAMB-BE/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerMasterSupplier struct {
	Service services.ServiceMasterSupplier
}

func InitControllerMasterSupplier(g *gin.Engine, service services.ServiceMasterSupplier) {
	handler := &ControllerMasterSupplier{
		Service: service,
	}

	routeAPI := g.Group("/api/v1/")
	routeAPI.GET("/suppliers", handler.ListSupplierController)
}

// ListSupplierController godoc
// @Tags MasterSupplier
// @Summary ListSupplierController
// @Description ListSupplierController
// @ID ListSupplierController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/suppliers [get]
func (h *ControllerMasterSupplier) ListSupplierController(c *gin.Context) {

	res, err := h.Service.ListMasterSupplierService()
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}

}
