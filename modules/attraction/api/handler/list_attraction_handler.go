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

// HandleListAttraction	godoc
// @Summary      List attraction
// @Description  List attraction
// @Tags         Attraction
// @Produce      json
// @Param        FilterAttractionReq  query	req.FilterAttractionReq  true  "FilterAttractionReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /attraction/list [get]
func (h attractionHandler) HandleListAttraction(c *gin.Context) {
	// Bind request
	var filter req.FilterAttractionReq
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.JsonLogger.Error("HandleListAttraction.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Create attraction
	attractions, err := h.listAttractionUseCase.ExecuteListAttraction(
		c.Request.Context(),
		mapper.ConvertFilterAttractionReqToEntity(filter),
	)
	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", attractions))
}
