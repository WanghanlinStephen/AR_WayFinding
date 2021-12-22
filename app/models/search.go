package models

type SearchInput struct {
	Source string
	Destination string
}

type SearchOutput struct {
	Angle 	float64
}