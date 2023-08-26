package controllers

import (
	services "SAMB-BE/services/master_product"
	"SAMB-BE/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerMasterProduct struct {
	Service services.ServiceMasterProduct
}

func InitControllerMasterProduct(g *gin.Engine, service services.ServiceMasterProduct) {
	handler := &ControllerMasterProduct{
		Service: service,
	}

	routeAPI := g.Group("/api/v1/")
	routeAPI.GET("/products", handler.ListProductController)
}

// ListProductController godoc
// @Tags MasterProduct
// @Summary ListProductController
// @Description ListProductController
// @ID ListProductController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/products [get]
func (h *ControllerMasterProduct) ListProductController(c *gin.Context) {

	res, err := h.Service.ListMasterProductService()
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}

}
