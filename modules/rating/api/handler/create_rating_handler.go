package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/rating/api/mapper"
	"trekkstay/modules/rating/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateRating	godoc
// @Summary      Create rating
// @Description  Create rating, requires authentication with user role
// @Tags         Rating
// @Produce      json
// @Param        CreateRatingReq  body	req.CreateRatingReq  true  "CreateRatingReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /rating/create [post]
// @Security     JWT
func (h ratingHandler) HandleCreateRating(c *gin.Context) {
	// Bind request
	var createRatingRequest req.CreateRatingReq
	if err := c.ShouldBindJSON(&createRatingRequest); err != nil {
		log.JsonLogger.Error("HandleCreateRating.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Call service
	err := h.createRatingUseCase.ExecuteCreateRating(ctx,
		mapper.ConvertCreateRatingReqToEntity(createRatingRequest),
	)

	if err != nil {
		panic(err)
	}

	// Response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
