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
	Weight int
}
