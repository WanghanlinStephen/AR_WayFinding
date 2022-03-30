package models
type DeleteNodeInput struct {
	Id int
}

type DeleteConnectionByIDInput struct {
	Id int
}

type DeleteConnectionByNodeInput struct {
	SourceId int
	DestinationId int
}

type DeleteInput struct {
	Id int
}

type DeleteMapByNameAndFloorInput struct {
	Name string
	Floor int
}