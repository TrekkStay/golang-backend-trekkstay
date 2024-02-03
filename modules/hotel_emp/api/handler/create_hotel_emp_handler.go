package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"trekkstay/core"
	res "trekkstay/core/response"
	"trekkstay/modules/hotel_emp/api/mapper"
	"trekkstay/modules/hotel_emp/api/model/req"
	"trekkstay/pkgs/log"
)

// HandleCreateHotelEmp	godoc
// @Summary      Create new hotel employee account
// @Description  Create new hotel employee account, require hotel owner permission and hotel profile already created
// @Tags         Hotel Employee
// @Produce      json
// @Param        CreateHotelEmpReq  body	req.CreateHotelEmpReq  true  "CreateHotelEmpReq JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-emp/create-emp [post]
// @Security 	JWT
func (h hotelEmpHandler) HandleCreateHotelEmp(c *gin.Context) {
	// Bind request
	var createHotelEmpReq req.CreateHotelEmpReq
	if err := c.ShouldBindJSON(&createHotelEmpReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelEmp.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&createHotelEmpReq); err != nil {
		log.JsonLogger.Error("HandleCreateHotelEmp.validate_request",
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

				if e.Field() == "Contract" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid contract")))
				}

				if e.Field() == "BaseSalary" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid base salary")))
				}
			}

			// If no field matched, return default error
			panic(res.ErrFieldValidationFailed(err))
		}
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Execute use case
	if err := h.createHotelEmpUseCase.ExecuteCreateHotelEmp(
		ctx,
		mapper.ConvertCreateHotelEmpReqEntity(createHotelEmpReq),
	); err != nil {
		log.JsonLogger.Error("HandleCreateHotelOwner.exec_create_hotel_owner",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusCreated, "success", nil))
}
