package handler

import (
	"github.com/gin-gonic/gin"
	res "trekkstay/core/response"
)

// HandleListProvince godoc
// @Summary      List provinces
// @Description  List all provinces
// @Tags         Region
// @Produce      json
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /region/list-province [get]
func (h regionHandler) HandleListProvince(c *gin.Context) {
	provinces, err := h.listProvinceUseCase.ExecuteListProvince(c.Request.Context())

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(200, "Success", provinces))
}
