package covid19

import "time"

type (
	GetCovid19Request struct {
		Country string    `json:"country"`
		Status  string    `json:"status"`
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
	}
	GetCovid19Response struct {
		Data
	}
)
