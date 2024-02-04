package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel/api/mapper"
	"trekkstay/modules/hotel/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreatHotel	godoc
// @Summary      Create new hotel
// @Description  Create new hotel, requires authentication with owner role
// @Tags         Hotel
// @Produce      json
// @Param        CreateHotelReq  body	req.CreateHotelReq  true  "CreateHotelReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/create [post]
// @Security     JWT
func (h hotelHandler) HandleCreatHotel(c *gin.Context) {
	// Bind request
	var createHotelReq req.CreateHotelReq
	if err := c.ShouldBindJSON(&createHotelReq); err != nil {
		log.JsonLogger.Error("HandleCreatHotel.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createHotelReq); err != nil {
		log.JsonLogger.Error("HandleCreatHotel.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Create hotel
	if err := h.createHotelUseCase.ExecuteCreateHotel(
		ctx,
		mapper.ConvertCreateHotelReqToEntity(createHotelReq),
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
