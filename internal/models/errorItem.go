package models

type ErrorItem struct {
	Message   string `json:"errorMessage"`
	ErrorItem error  `json:"errorItem"`
	Code      int    `json:"code"`
}
