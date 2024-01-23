package core

const CurrentRequesterKeyString = ""

type CurrentRequesterKeyStruct struct{}

type Requester interface {
	GetUserId() string
}

var _ Requester = (*RestRequester)(nil)

type RestRequester struct {
	Id string
}

func (u RestRequester) GetUserId() string {
	return u.Id
}
