package core

const CurrentRequesterKey = ""

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
