package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel/api/mapper"
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleFilterHotel	godoc
// @Summary      Filter hotel
// @Description  Filter hotel
// @Tags         Hotel
// @Produce      json
// @Param        FilterHotelReq  query	req.FilterHotelReq  true  "FilterHotelReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/filter [get]
// @Security     JWT
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

	var hotels *core.Pagination

	cacheKey := c.Request.URL.RequestURI()
	err := h.cache.Get(cacheKey, &hotels)
	if err == nil {
		log.JsonLogger.Debug("cache hit", slog.String("key", cacheKey))
		res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotels))
		return
	}

	hotels, err = h.filterHotelUseCase.FilterHotel(
		c.Request.Context(),
		mapper.ConvertFilterHotelReqToEntity(filterHotelReq),
		filterHotelReq.Page,
		filterHotelReq.Limit,
	)

	if err != nil {
		panic(err)
	}

	_ = h.cache.SetWithExpiration(cacheKey, hotels, 1*time.Minute)

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotels))
}
