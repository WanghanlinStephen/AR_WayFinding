package models


type Node struct {
	Id int
	NameEnglish string
	NameChinese string
	NameTraditionalChinese string
	Latitude float64
	Longitude float64
	IntersectionalAngle float64
	IsStaircase int
}


type Connection struct {
	Id int
	Source Node
	Destination Node
	Time float64
	MapId int
}

type GetNodeOutput struct {
	Node Node
}

type GetNodesOutput struct {
	Nodes []Node
}
type GetNodesByMapId struct {
	Id int
}
type GetNodesByMapIdOutput struct {
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
	Floor int
}
type GetMapsOutput struct {
	Maps []Map
}
type GetMapByIdInput struct {
	Id int
}
type GetMapByNameInput struct {
	Name string
}
type GetMapByIdOutput struct {
	Map Map
}
type GetMapByNameOutput struct {
	Map []Map
}
type GetMapNamesOutput struct {
	Names []string
}
type GetMapIdByNodeIdInput struct {
	Id int
}
type GetMapIdByNodeIdOutput struct {
	Id int
}