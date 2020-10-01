package entity

type Sensor struct {
	Uuid 			string `json:"uuid" gorm:"primary_key"`
	MaxLast30Days 	int `json:"maxLast30Days"`
	AvgLast30Days   int `json:"avgLast30Days"`
	StatusExceeded	int `json:"status_exceeded"`
	StatusOk		int `json:"status_ok"`
}
