package strategy

import (
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"pro/app/model"
	"sync"
)

var CyberPortMap *Graph

type node struct {
	id                   string
	nameEnglish          string
	nameChinese          string
	nameChineseTradition string
	Latitude             string
	Longitude            string
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
	return n.id
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
	g.NodeMap[nodeID1.id] = nodeID1
	g.NodeMap[nodeID2.id] = nodeID2

	if _, ok := g.edge[nodeID1.id]; ok {
		g.edge[nodeID1.id][nodeID2.id] = w
	} else {
		tempMap := make(map[string]float64)
		tempMap[nodeID2.id] = w
		g.edge[nodeID1.id] = tempMap
	}
}

//restore data from database and initialize graph.
func Initialization() {
	g := NewGraph()
	//fixme:修复
	connectionsList,_:=model.GetConnectionsList()
	for _,value := range connectionsList{
		source:=node{
			id:                   fmt.Sprint(value.Source.Id),
			nameEnglish:          value.Source.NameEnglish,
			nameChinese:          value.Source.NameChinese,
			nameChineseTradition: value.Source.NameTraditionalChinese,
			Latitude:			  fmt.Sprint(value.Source.Latitude),
			Longitude:			  fmt.Sprint(value.Source.Longitude),
		}
		destination:=node{
			id:                   fmt.Sprint(value.Destination.Id),
			nameEnglish:          value.Destination.NameEnglish,
			nameChinese:          value.Destination.NameChinese,
			nameChineseTradition: value.Destination.NameTraditionalChinese,
			Latitude:			  fmt.Sprint(value.Destination.Latitude),
			Longitude:			  fmt.Sprint(value.Destination.Longitude),
		}
		g.AddEdge(source,destination,float64(value.Time))
	}
	CyberPortMap =g

	//todo:test model
	source:=g.NodeMap["1"]
	destination:=g.NodeMap["3"]
	//todo:test next step
	shortestDistance,nextStep:=CyberPortMap.Dijkstra(source.id,destination.id)
	//todo:test direction
	direction,angle:=GetAngle(source.Longitude,source.Latitude,destination.Longitude,destination.Latitude)
	fmt.Printf("Source:%s to Destination%s with next step %s with a total weight %f,with a direction of %s, with an angle of %f",source.id,destination.id,nextStep,shortestDistance,direction, angle)


}