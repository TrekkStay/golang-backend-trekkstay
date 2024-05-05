package entity

import "trekkstay/core"

type UserEntity struct {
	core.BaseEntity `json:",inline"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}

func (UserEntity) TableName() string {
	return "users"
}
