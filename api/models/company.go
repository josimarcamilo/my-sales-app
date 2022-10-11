package models

type Company struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

var Companys []Company
