package handler

import (
	"github.com/gin-gonic/gin"
	res "trekkstay/core/response"
)

// HandleListDistrict godoc
// @Summary      List districts
// @Description  List all districts of a province
// @Tags         Region
// @Produce      json
// @Param        province_code query string true "Province code"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /region/list-district [get]
func (h regionHandler) HandleListDistrict(c *gin.Context) {
	provinceCode := c.Query("province_code")
	districts, err := h.listDistrictUseCase.ExecuteListDistrict(c.Request.Context(), provinceCode)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(200, "Success", districts))
}
