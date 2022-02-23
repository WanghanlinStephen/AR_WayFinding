package strategy

import (
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"pro/app/model"
	"pro/app/models"
	"sync"
)

var CyberPortMap *Graph
var ConnectionsList []models.Connection

type node struct {
	Id                   string
	NameEnglish          string
	NameChinese          string
	NameChineseTradition string
	Latitude             string
	Longitude            string
	//the angle between wall and horizontal direction
	IntersectionalAngle  float64
}

type Edge struct {
	src    Node
	dst    Node
	weight float64
}

type Graph struct {
	sync.RWMutex
	edge    map[string]map[string]float64
	NodeMap map[string]node // record all the node in a graph
}

type Node interface {
	NodeID() string
}

func (n *node) NodeID() string {
	return n.Id
}

func NewEdge(src Node, dst Node, w float64) *Edge {
	return &Edge{src: src, dst: dst, weight: w}
}

func NewGraph() *Graph {
	return &Graph{
		edge:    make(map[string]map[string]float64),
		NodeMap: make(map[string]node),
	}
}

func (g *Graph) AddEdge(nodeID1 node, nodeID2 node, w float64) {
	g.Lock()
	defer g.Unlock()

	if nodeID1 == nodeID2 {
		panic("can't add same vertex in one edge")
		return
	}

	if w == 0 {
		panic("weight can't use 0")
		return
	}

	// record each vertex
	g.NodeMap[nodeID1.Id] = nodeID1
	g.NodeMap[nodeID2.Id] = nodeID2

	if _, ok := g.edge[nodeID1.Id]; ok {
		g.edge[nodeID1.Id][nodeID2.Id] = w
	} else {
		tempMap := make(map[string]float64)
		tempMap[nodeID2.Id] = w
		g.edge[nodeID1.Id] = tempMap
	}
}

//restore data from database and initialize graph.
func Initialization() {
	g := NewGraph()
	//fixme:修复
	ConnectionsList,_=model.GetConnectionsList()
	value:=ConnectionsList;
	fmt.Println(value);
	for _,value := range ConnectionsList{
		source:=node{
			Id:                   fmt.Sprint(value.Source.Id),
			NameEnglish:          value.Source.NameEnglish,
			NameChinese:          value.Source.NameChinese,
			NameChineseTradition: value.Source.NameTraditionalChinese,
			Latitude:			  fmt.Sprint(value.Source.Latitude),
			Longitude:			  fmt.Sprint(value.Source.Longitude),
			IntersectionalAngle:  value.Source.IntersectionalAngle,
		}
		destination:=node{
			Id:                   fmt.Sprint(value.Destination.Id),
			NameEnglish:          value.Destination.NameEnglish,
			NameChinese:          value.Destination.NameChinese,
			NameChineseTradition: value.Destination.NameTraditionalChinese,
			Latitude:			  fmt.Sprint(value.Destination.Latitude),
			Longitude:			  fmt.Sprint(value.Destination.Longitude),
			IntersectionalAngle:  value.Destination.IntersectionalAngle,
		}
		g.AddEdge(source,destination,float64(value.Time))
	}
	CyberPortMap =g
}