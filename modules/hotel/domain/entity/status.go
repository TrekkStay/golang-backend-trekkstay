package entity

type Status int

const (
	INACTIVE Status = iota
	ACTIVE
	UNVERIFIED
)

func (status Status) Value() string {
	switch status {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	case UNVERIFIED:
		return "unverified"
	default:
		return "-"
	}
}
