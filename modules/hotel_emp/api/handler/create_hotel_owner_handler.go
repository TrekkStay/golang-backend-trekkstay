package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel_emp/api/mapper"
	"trekkstay/modules/hotel_emp/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateHotelOwner	godoc
// @Summary      Create new hotel owner account
// @Description  Create new hotel owner account
// @Tags         Hotel Employee
// @Produce      json
// @Param        CreateHotelOwnerReq  body	req.CreateHotelOwnerReq  true  "CreateHotelOwnerReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-emp/create-owner [post]
func (h hotelEmpHandler) HandleCreateHotelOwner(c *gin.Context) {
	// Bind request
	var createHotelOwnerReq req.CreateHotelOwnerReq
	if err := c.ShouldBindJSON(&createHotelOwnerReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelOwner.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createHotelOwnerReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelOwner.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				if e.Field() == "Email" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid email")))
				}

				if e.Field() == "Phone" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid phone number")))
				}

				if e.Field() == "Password" {
					panic(res.ErrFieldValidationFailed(errors.New("password too weak")))
				}
			}

			// If no field matched, return default error
			panic(res.ErrFieldValidationFailed(err))
		}
	}

	// Execute use case
	if err := h.createHotelOwnerUseCase.ExecuteCreateHotelOwner(
		c.Request.Context(),
		mapper.ConvertCreateHotelOwnerReqEntity(createHotelOwnerReq),
	); err != nil {
		log.JsonLogger.Error("HandleCreateHotelOwner.exec_create_hotel_owner",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
