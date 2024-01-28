package core

const CurrentRequesterKeyString = ""

type CurrentRequesterKeyStruct struct{}

type Requester interface {
	GetUserID() string
}

var _ Requester = (*RestRequester)(nil)

type RestRequester struct {
	ID string
}

func (u RestRequester) GetUserID() string {
	return u.ID
}
