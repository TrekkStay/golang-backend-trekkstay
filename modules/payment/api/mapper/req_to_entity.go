package mapper

import (
	"trekkstay/modules/payment/api/model/req"
	"trekkstay/modules/payment/domain/entity"
)

func ConvertCreatePaymentReqToEntity(req req.CreatePaymentReq) entity.PaymentEntity {
	return entity.PaymentEntity{
		ReservationID: req.ReservationID,
		Amount:        req.Amount,
	}
}
