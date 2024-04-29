package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel/api/mapper"
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleSearchHotel	godoc
// @Summary      Search hotel
// @Description  Search hotel
// @Tags         Hotel
// @Produce      json
// @Param        SearchHotelReq  query	req.SearchHotelReq  true  "SearchHotelReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/search [get]
func (h hotelHandler) HandleSearchHotel(c *gin.Context) {
	// Bind request
	var searchHotelReq req.SearchHotelReq
	if err := c.ShouldBindQuery(&searchHotelReq); err != nil {
		log.JsonLogger.Error("HandleSearchHotel.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	var hotels *core.Pagination

	// Cache
	//cacheKey := c.Request.URL.RequestURI()
	//if err := h.cache.Get(cacheKey, &hotels); err == nil {
	//	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotels))
	//	return
	//}

	hotels, err := h.searchHotelUseCase.ExecuteSearchHotel(
		c.Request.Context(),
		mapper.ConvertSearchHotelReqToEntity(searchHotelReq),
		searchHotelReq.Page,
		searchHotelReq.Limit,
	)

	if err != nil {
		panic(err)
	}

	// Set cache
	//_ = h.cache.SetWithExpiration(cacheKey, hotels, 1*time.Minute)

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", hotels))
}
