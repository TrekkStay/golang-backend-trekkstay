package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel/api/mapper"
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/pkgs/log"
)

func (h hotelHandler) HandleFilterHotel(c *gin.Context) {
	// Bind request
	var filterHotelReq req.FilterHotelReq
	if err := c.ShouldBindQuery(&filterHotelReq); err != nil {
		log.JsonLogger.Error("HandleCreatHotel.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	hotels, err := h.filterHotelUseCase.FilterHotel(
		c.Request.Context(),
		mapper.ConvertFilterHotelReqToEntity(filterHotelReq),
		filterHotelReq.Page,
		filterHotelReq.Limit,
	)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotels))
}
