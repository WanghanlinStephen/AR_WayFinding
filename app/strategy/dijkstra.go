package strategy

import (
	"fmt"
	"math"
)

func reverseArrays(input []string)[]string{
	if len(input) == 0 {
		return input
	}
	return append(reverseArrays(input[1:]), input[0])
}

//fixme:未获取数据
func (g *Graph) Dijkstra(src string, dst string, mapId int) (shortDis float64,lastPointer string,fullPathway []string) {

	// path:store previous node
	path:=make(map[string]string)
	// pathList: store previous nodes
	pathList:=make(map[string][]string)
	// infinity
	infinity := math.Inf(1)

	// init a short distance map to record the shortest distance from src
	distance := make(map[string]float64)
	for nodeID := range g.NodeMap {
		if nodeID == src {
			distance[nodeID] = 0
		} else {
			distance[nodeID] = infinity
		}
	}

	q := NewQueue()
	q.Push(src)

	for !q.Empty() {
		v := q.Pop()
		e, ok := v.(string)
		if !ok {
			return 0,"",nil
		}

		for nodeID := range g.NodeMap {
			//check map Id
			if g.NodeMap[nodeID].MapId != mapId {
				continue
			}
			if nodeID == e {
				continue
			}
			//Current Node (0,A) + shortest distance (A,B) < shortest distance (0,B) the update
			if distance[e] + g.edge[e][nodeID] < distance[nodeID] && g.edge[e][nodeID] != 0 {
				//https://blog.csdn.net/jinixin/article/details/52247763
				distance[nodeID] = g.edge[e][nodeID] + distance[e]
				q.Push(nodeID)
				//path[B]=A means the shortest path of OB composed of OA+AB
				path[nodeID]=e
			}
		}
	}

	//distance array
	for nodeID := range g.NodeMap {
		//check map Id
		if g.NodeMap[nodeID].MapId != mapId {
			continue
		}
		temp:=nodeID
		for distance[temp]!=0 {
			pathList[nodeID]=append(pathList[nodeID],temp)
			temp=path[temp]
		}
		pathList[nodeID]=append(pathList[nodeID],temp)
		pathList[nodeID]=reverseArrays(pathList[nodeID])
	}
	// strategy layer : only return the next close node
	var nextStep string
	if len(pathList[dst])>1{
		nextStep=pathList[dst][1]
	}else{
		nextStep="-1"
	}
	fullPath:=pathList[dst]
	fmt.Println(fullPath)
	return distance[dst],nextStep,fullPath
}
