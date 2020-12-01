package models

import (
	"time"
)

type Profile struct {
	IDProfile int    `json:"idProfile"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Type      string `json:"type"`
	Time      time.Time
}

type AllProfile []Profile
