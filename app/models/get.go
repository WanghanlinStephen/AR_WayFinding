package models


type Node struct {
	Id int
	NameEnglish string
	NameChinese string
	NameTraditionalChinese string
	Latitude float64
	Longitude float64
	IntersectionalAngle float64
}


type Connection struct {
	Id int
	Source Node
	Destination Node
	Time int
}

type GetNodeOutput struct {
	Node Node
}

type GetNodesOutput struct {
	Nodes []Node
}

type GetConnectionOutput struct {
	Connection Connection

}

type GetConnectionsOutput struct {
	Connections []Connection
}
type Map struct {
	Id int
	Name string
	Url string
}
type GetMapsOutput struct {
	Maps []Map
}
type GetMapByNameInput struct {
	Name string
}
type GetMapByNameOutput struct {
	Map Map
}