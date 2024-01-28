package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/user/api/mapper"
	"trekkstay/modules/user/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleUpdateUser	godoc
// @Summary      Update user
// @Description  Update user
// @Tags         User
// @Produce      json
// @Param        UpdateUserReq  body	req.UpdateUserReq  true  "UpdateUserReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/update [patch]
// @Security     JWT
func (h *userHandler) HandleUpdateUser(c *gin.Context) {
	var updateUserReq req.UpdateUserReq
	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		log.JsonLogger.Error("HandleUpdateUser.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Update user
	err := h.updateUserUseCase.ExecUpdateUser(ctx, mapper.ConvertUpdateUserReqToUserEntity(updateUserReq))

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
