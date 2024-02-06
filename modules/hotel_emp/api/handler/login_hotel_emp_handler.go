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

// HandleLoginHotelEmp	godoc
// @Summary      Login hotel employee
// @Description  Login hotel employee by email and password
// @Tags         Hotel Employee
// @Produce      json
// @Param        LoginHotelEmpReq  body	req.LoginHotelEmpReq  true  "LoginHotelEmpReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /hotel-emp/login [post]
func (h hotelEmpHandler) HandleLoginHotelEmp(c *gin.Context) {
	// Bind request
	var loginHotelEmpReq req.LoginHotelEmpReq
	if err := c.ShouldBindJSON(&loginHotelEmpReq); err != nil {
		log.JsonLogger.Error("HandleLoginHotelEmp.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&loginHotelEmpReq); err != nil {
		log.JsonLogger.Error("HandleLoginHotelEmp.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				if e.Field() == "Email" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid email")))
				}
			}

			// If no field matched, return default error
			panic(res.ErrFieldValidationFailed(err))
		}
	}

	// Login hotel employee
	hotelEmp, err := h.loginHotelEmpUseCase.ExecuteLoginHotelEmp(c.Request.Context(),
		mapper.ConvertLoginHotelEmpReqEntity(loginHotelEmpReq))
	if err != nil {
		panic(err)
	}

	// Convert response
	empResponse := mapper.CovertUserEntityToLoginHotelEmpRes(*hotelEmp)

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", empResponse))
}
