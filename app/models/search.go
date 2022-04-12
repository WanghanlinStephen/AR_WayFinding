package models

type SearchInput struct {
	Source string
	Destination string
}

type SearchOutput struct {
	Angle 	float64
	IsSameFloor bool
	Floor int
}

type FetchPathInput struct {
	Source string
	Destination string
}

type FetchPathOutput struct {
	Path []Node
	IsSameFloor bool
	DestinationId string
	Floor int
}
