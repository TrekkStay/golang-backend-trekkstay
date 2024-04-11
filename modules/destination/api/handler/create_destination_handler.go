package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/destination/api/mapper"
	"trekkstay/modules/destination/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateDestination	godoc
// @Summary      Create new destination
// @Description  Create new destination
// @Tags         Destination
// @Produce      json
// @Param        CreateDestinationReq  body	req.CreateDestinationReq  true  "CreateDestinationReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /destination/create [post]
// @Security     JWT
func (h destinationHandler) HandleCreateDestination(c *gin.Context) {
	var createDestinationRequest req.CreateDestinationReq

	if err := c.ShouldBindJSON(&createDestinationRequest); err != nil {
		log.JsonLogger.Error("HandleCreateDestination.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createDestinationRequest); err != nil {
		log.JsonLogger.Error("HandleCreateDestination.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Create destination
	if err := h.createDestinationUseCase.ExecuteCreateDestination(
		c.Request.Context(),
		mapper.ConvertCreateDestinationReqToEntity(createDestinationRequest),
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
