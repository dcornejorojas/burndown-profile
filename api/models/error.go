package models

type Error struct {
	Flag    bool    `json:"flag"`
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}
