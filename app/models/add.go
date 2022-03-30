package models
type AddNodeInput struct {
	Id int
	NameEnglish string
	NameChinese string
	NameTraditionalChinese string
	Latitude float64
	Longitude float64
	IntersectionalAngle float64
}

type AddConnectionInput struct {
	SourceId int
	DestinationId int
	Weight float64
	MapId int
}

type AddStaircaseInput struct {
	Latitude float64
	Longitude float64
	MapId int
}

type AddMapInput struct {
	Url string
	Name string
	Floor int
}