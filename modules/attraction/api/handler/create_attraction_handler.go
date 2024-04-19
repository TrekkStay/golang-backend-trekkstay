package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/attraction/api/mapper"
	"trekkstay/modules/attraction/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateAttraction	godoc
// @Summary      Create new attraction
// @Description  Create new attraction
// @Tags         Attraction
// @Produce      json
// @Param        CreateAttractionReq  body	req.CreateAttractionReq  true  "CreateAttractionReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /attraction/create [post]
// @Security     JWT
func (h attractionHandler) HandleCreateAttraction(c *gin.Context) {
	// Bind request
	var createAttractionReq req.CreateAttractionReq
	if err := c.ShouldBindJSON(&createAttractionReq); err != nil {
		log.JsonLogger.Error("HandleCreateAttraction.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createAttractionReq); err != nil {
		log.JsonLogger.Error("CreateAttractionReq.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrFieldValidationFailed(err))
	}

	// Create attraction
	if err := h.createAttractionUseCase.ExecuteCreateAttraction(
		c.Request.Context(),
		mapper.ConvertCreateAttractionReqToEntity(createAttractionReq),
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
