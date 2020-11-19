package models

type Image struct {
	IDImage int64  `json:"idImage"`
	Path    string `json:"path"`
	Name    string `json:"name"`
}

type Images []Image
