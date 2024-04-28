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

// HandleCreateReservation	godoc
// @Summary      Create new reservation
// @Description  Create new reservation, requires authentication with user role
// @Tags         Reservation
// @Produce      json
// @Param        CreateReservationReq  body	req.CreateReservationReq  true  "CreateReservationReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /reservation/create [post]
// @Security     JWT
func (h reservationHandler) HandleCreateReservation(c *gin.Context) {
	// Bind request
	var createReservationRequest req.CreateReservationReq
	if err := c.ShouldBindJSON(&createReservationRequest); err != nil {
		log.JsonLogger.Error("HandleCreateReservation.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	entity := mapper.ConvertCreateReservationReqToEntity(createReservationRequest)
	reservation, err := h.createReservationUseCase.ExecuteCreateReservation(
		ctx,
		&entity,
	)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", reservation))
}
