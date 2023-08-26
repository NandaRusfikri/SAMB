package controllers

import (
	"SAMB-BE/schemas"
	services "SAMB-BE/services/barang_masuk"
	"SAMB-BE/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerPenerimaanBarang struct {
	Service services.ServicePenerimaanBarang
}

func InitControllerPenerimaanBarang(g *gin.Engine, service services.ServicePenerimaanBarang) {
	handler := &ControllerPenerimaanBarang{
		Service: service,
	}

	routeAPI := g.Group("/api/v1/barang_masuk")
	routeAPI.POST("/header", handler.CreatePenerimaanBarangHeaderController)
	routeAPI.POST("/detail", handler.CreatePenerimaanBarangDetailController)
	routeAPI.GET("/header", handler.ListBarangMasukHeaderController)
	routeAPI.GET("/detail/:TrxInPK", handler.ListBarangMasukDetailController)
	routeAPI.GET("/header/:TrxInPK", handler.GetByIdBarangMasukHeaderController)

}

// CreatePenerimaanBarangHeaderController godoc
// @Tags PenerimaanBarang
// @Summary Create Penerimaan Barang Header
// @Description CreatePenerimaanBarangHeaderController
// @ID CreatePenerimaanBarangHeaderController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.PenerimaanBarangHeaderRequest true "body data"
// @Success 200
// @Router /v1/barang_masuk/header [post]
func (h *ControllerPenerimaanBarang) CreatePenerimaanBarangHeaderController(c *gin.Context) {
	var input schemas.PenerimaanBarangHeaderRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		util.APIResponse(c, "request invalid", http.StatusBadRequest, 0, nil)
		return
	}

	fmt.Printf("%+v\n", input)

	res, err := h.Service.CreatePenerimaanBarangHeaderService(input)
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}

}

// CreatePenerimaanBarangDetailController godoc
// @Tags PenerimaanBarang
// @Summary Create Penerimaan Barang Detail
// @Description CreatePenerimaanBarangDetailController
// @ID CreatePenerimaanBarangDetailController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.PenerimaanBarangDetailRequest true "body data"
// @Success 200
// @Router /v1/barang_masuk/detail [post]
func (h *ControllerPenerimaanBarang) CreatePenerimaanBarangDetailController(c *gin.Context) {
	var input schemas.PenerimaanBarangDetailRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		util.APIResponse(c, "request invalid", http.StatusBadRequest, 0, nil)
		return
	}

	res, err := h.Service.CreatePenerimaanBarangDetailService(input)
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}
}

// ListBarangMasukHeaderController godoc
// @Tags PenerimaanBarang
// @Summary ListBarangMasukHeaderController
// @Description ListBarangMasukHeaderController
// @ID ListBarangMasukHeaderController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/barang_masuk/header [get]
func (h *ControllerPenerimaanBarang) ListBarangMasukHeaderController(c *gin.Context) {

	res, err := h.Service.ListBarangMasukHeaderService()
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}
}

// ListBarangMasukDetailController godoc
// @Tags PenerimaanBarang
// @Summary ListBarangMasukDetailController
// @Description ListBarangMasukDetailController
// @ID ListBarangMasukDetailController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param TrxInPK path int true "TrxInPK"
// @Success 200
// @Router /v1/barang_masuk/detail/{TrxInPK} [get]
func (h *ControllerPenerimaanBarang) ListBarangMasukDetailController(c *gin.Context) {

	ticket_id := c.Param("TrxInPK")

	convert_ticket_id, errParse := strconv.ParseUint(ticket_id, 10, 64)

	if errParse != nil {
		util.APIResponse(c, "request invalid", http.StatusBadRequest, 0, nil)
		return
	}

	res, err := h.Service.ListBarangMasukDetailService(uint(convert_ticket_id))
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}
}

// GetByIdBarangMasukHeaderController godoc
// @Tags PenerimaanBarang
// @Summary GetByIdBarangMasukHeaderController
// @Description GetByIdBarangMasukHeaderController
// @ID GetByIdBarangMasukHeaderController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param TrxInPK path int true "TrxInPK"
// @Success 200
// @Router /v1/barang_masuk/header/{TrxInPK} [get]
func (h *ControllerPenerimaanBarang) GetByIdBarangMasukHeaderController(c *gin.Context) {

	ticket_id := c.Param("TrxInPK")

	convert_ticket_id, errParse := strconv.ParseUint(ticket_id, 10, 64)

	if errParse != nil {
		util.APIResponse(c, "request invalid", http.StatusBadRequest, 0, nil)
		return
	}

	res, err := h.Service.GetByIdBarangMasukHeaderService(uint(convert_ticket_id))
	if err.Error != nil {
		util.APIResponse(c, err.Error.Error(), err.Code, 0, nil)
		return
	} else {
		util.APIResponse(c, "success", http.StatusOK, 0, res)
		return
	}
}
