package entity

type HotelEmpStatus int

const (
	ACTIVE HotelEmpStatus = iota
	BLOCKED
	UNVERIFIED
)

func (status HotelEmpStatus) Value() string {
	switch status {
	case ACTIVE:
		return "ACTIVE"
	case BLOCKED:
		return "BLOCKED"
	case UNVERIFIED:
		return "UNVERIFIED"
	}

	return ""
}
