package model

import (
	"time"
)

type User struct {
	Id                               string    `json:"id" gorm:"primaryKey"`
	Name                             string    `json:"name" gorm:"unique"`
	Password                         string    `json:"password"`
	SalesforceLoginDomain            string    `json:"salesforceLoginDomain"`
	SalesforceMyDomain               string    `json:"salesforceMyDomain"`
	SalesforceClientId               string    `json:"salesforceClientId"`
	SalesforceClientSecret           string    `json:"salesforceClientSecret"`
	SalesforceUserName               string    `json:"salesforceUserName"`
	SalesforcePassword               string    `json:"salesforcePassword"`
	SalesforceAccessTokenLastUpdated time.Time `json:"salesforceAccessTokenLastUpdated"`
	UpdatedAt                        time.Time `json:"updatedAt"`
	CreatedAt                        time.Time `json:"createdAt"`
	DeletedAt                        time.Time `json:"deletedAt"`
}

type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
