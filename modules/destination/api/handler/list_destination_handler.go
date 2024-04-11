package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
)

// HandleListDestination	godoc
// @Summary      List destination
// @Description  List destination
// @Tags         Destination
// @Produce      json
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /destination/list [get]
func (h destinationHandler) HandleListDestination(c *gin.Context) {
	destinations, err := h.listDestinationUseCase.ExecuteListDestination(c)

	if err != nil {
		panic(err)
		return
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", destinations))
}
