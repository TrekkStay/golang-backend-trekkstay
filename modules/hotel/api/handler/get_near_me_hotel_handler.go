package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleGetNearMeHotel	godoc
// @Summary      Get near me hotel
// @Description  Get near me hotel
// @Tags         Hotel
// @Produce      json
// @Param        GetNearMeHotelReq  query	req.GetNearMeHotelReq  true  "GetNearMeHotelReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/filter/near-me [get]
func (h hotelHandler) HandleGetNearMeHotel(c *gin.Context) {
	// Bind request
	var getNearMeHotelReq req.GetNearMeHotelReq
	if err := c.ShouldBind(&getNearMeHotelReq); err != nil {
		log.JsonLogger.Error("HandleGetNearMeHotel.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	hotels, err := h.getNearMeHotelUseCase.ExecGetNearMeHotelUseCase(
		c.Request.Context(),
		getNearMeHotelReq.Lat,
		getNearMeHotelReq.Lng,
		getNearMeHotelReq.MaxDistance,
	)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(200, "success", hotels))
}
