package entity

type PaymentStatus int

const (
	SUCCESS PaymentStatus = iota + 1
	FAILED
	PENDING
)

func (status PaymentStatus) Value() string {
	switch status {
	case SUCCESS:
		return "SUCCESS"
	case FAILED:
		return "FAILED"
	case PENDING:
		return "PENDING"
	}

	return ""
}
