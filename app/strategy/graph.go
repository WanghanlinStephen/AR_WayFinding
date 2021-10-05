package strategy

import (
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"pro/app/model"
	"sync"
)

type node struct {
	id string
	nameEnglish string
	nameChinese string
	nameChineseTradition string
	latitude string
	longitude string
}

type Edge struct {
	src    Node
	dst    Node
	weight float64
}

type Graph struct {
	sync.RWMutex
	edge    map[string]map[string]float64
	nodeMap map[string]node // record all the node in a graph
}

type Node interface {
	NodeID() string
}

func NewNode(id string) Node {
	return &node{id: id}
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
		nodeMap: make(map[string]node),
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
	g.nodeMap[nodeID1.id] = nodeID1
	g.nodeMap[nodeID2.id] = nodeID2

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
	//fixme:测试后删除
	//config.Run()
	//if err := model.Run(); err != nil {
	//	fmt.Println("数据库链接失败:", err)
	//	return
	//}
	g := NewGraph()
	//fixme:修复
	connectionsList,_:=model.GetConnectionsList()
	for _,value := range connectionsList{
		source:=node{
			id:                   fmt.Sprint(value.Source.Id),
			nameEnglish:          value.Source.NameEnglish,
		}
		destination:=node{
			id:                   fmt.Sprint(value.Destination.Id),
			nameEnglish:          value.Destination.NameEnglish,
		}
		g.AddEdge(source,destination,float64(value.Time))
	}
	fmt.Println(g)
	shortDis:=g.Dijkstra("1","2")
	fmt.Println("A->E shortest distance is:", shortDis)
}
//
//func main() {
//	g := NewGraph()
//	g.AddEdge("A", "B", 3)
//	g.AddEdge("A", "C", 2)
//	g.AddEdge("B", "E", 5)
//	g.AddEdge("B", "D", 2)
//	g.AddEdge("D", "E", -2)
//	g.AddEdge("C", "F", 1)
//	g.AddEdge("E", "F", 3)
//
//	fmt.Println(g)
//	shortDis := g.Dijkstra("A", "E")
//	fmt.Println("A->E shortest distance is:", shortDis)
//}
