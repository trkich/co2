package entity

import "time"

type Measurement struct {
	Uuid string `json:"uuid" gorm:"index"`
	Value int `json:"co2" gorm:"index"`
	Time  time.Time `json:"time" gorm:"index"`
}
