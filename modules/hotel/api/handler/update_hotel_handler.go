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

// HandleUpdateHotel	godoc
// @Summary      Update hotel
// @Description  Update hotel, requires authentication with owner role
// @Tags         Hotel
// @Produce      json
// @Param        UpdateHotelReq  body	req.UpdateHotelReq  true  "UpdateHotelReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel/update [patch]
// @Security     JWT
func (h hotelHandler) HandleUpdateHotel(c *gin.Context) {
	// Bind request
	var updateHotelReq req.UpdateHotelReq
	if err := c.ShouldBindJSON(&updateHotelReq); err != nil {
		log.JsonLogger.Error("HandleUpdateHotel.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&updateHotelReq); err != nil {
		log.JsonLogger.Error("HandleUpdateHotel.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	if err := h.updateHotelUseCase.ExecuteUpdateHotel(
		ctx,
		mapper.ConvertUpdateHotelReqToEntity(updateHotelReq)); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
