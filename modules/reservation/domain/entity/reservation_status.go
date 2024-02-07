package entity

type ReservationStatus int

const (
	UPCOMING ReservationStatus = iota + 1
	COMPLETED
	DONE
	CANCELLED
)

func (status ReservationStatus) Value() string {
	switch status {
	case UPCOMING:
		return "UPCOMING"
	case COMPLETED:
		return "COMPLETED"
	case DONE:
		return "DONE"
	case CANCELLED:
		return "CANCELLED"
	}

	return ""
}
