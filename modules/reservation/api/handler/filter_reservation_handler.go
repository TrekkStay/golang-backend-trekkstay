package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/reservation/api/mapper"
	"trekkstay/modules/reservation/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleFilterReservation	godoc
// @Summary      Filter reservation
// @Description  Filter reservation
// @Tags         Reservation
// @Produce      json
// @Param        FilterReservationReq  query	req.FilterReservationReq  true  "FilterReservationReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /reservation/filter [get]
// @Security     JWT
func (h reservationHandler) HandleFilterReservation(c *gin.Context) {
	// Bind request
	var filterReq req.FilterReservationReq
	if err := c.ShouldBindQuery(&filterReq); err != nil {
		log.JsonLogger.Error("HandleFilterReservation.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	pagination, err := h.filterReservationUseCase.ExecuteFilterReservation(
		ctx,
		mapper.ConvertFilterReservationReqToEntity(filterReq),
		filterReq.Page,
		filterReq.Limit,
	)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", pagination))
}
