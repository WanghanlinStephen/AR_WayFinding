package strategy

import (
	"fmt"
	"math"
)

func (g *Graph) Dijkstra(src string, dst string) (shortDis float64) {
	// infinity
	infinity := math.Inf(1)

	// init a short distance map to record the shortest distance from src
	distance := make(map[string]float64)
	for nodeID := range g.nodeMap {
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
			return 0
		}

		for nodeID := range g.nodeMap {
			if nodeID == e {
				continue
			}

			if g.edge[e][nodeID]+distance[e] < distance[nodeID] && g.edge[e][nodeID] != 0 {
				distance[nodeID] = g.edge[e][nodeID] + distance[e]
				q.Push(nodeID)
			}
		}
	}

	for node := range g.nodeMap {
		temp := fmt.Sprintf("from A -> %s =  %d", node, int(distance[node]))
		fmt.Println(temp)
	}
	return distance[dst]
}
