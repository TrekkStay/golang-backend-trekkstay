package handler

import (
	"github.com/gin-gonic/gin"
	res "trekkstay/core/response"
)

// HandleListWard godoc
// @Summary      List wards
// @Description  List all wards of a district
// @Tags         Region
// @Produce      json
// @Param        district_code query string true "District code"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /region/list-ward [get]
func (h regionHandler) HandleListWard(c *gin.Context) {
	districtCode := c.Query("district_code")
	wards, err := h.listWardUseCase.ExecuteListWard(c.Request.Context(), districtCode)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(200, "Success", wards))
}
