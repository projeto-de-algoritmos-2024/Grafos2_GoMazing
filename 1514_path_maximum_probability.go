import "container/heap"

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	graph := make([][]struct {
		to   int
		prob float64
	}, n)

	for i, edge := range edges {
		a, b := edge[0], edge[1]
		p := succProb[i]
		graph[a] = append(graph[a], struct {
			to   int
			prob float64
		}{b, p})
		graph[b] = append(graph[b], struct {
			to   int
			prob float64
		}{a, p})
	}

	maxProb := make([]float64, n)
	maxProb[start_node] = 1.0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start_node, prob: 1.0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		curNode, curProb := current.node, current.prob

		if curProb < maxProb[curNode] {
			continue
		}

		for _, neighbor := range graph[curNode] {
			nextNode, edgeProb := neighbor.to, neighbor.prob
			newProb := curProb * edgeProb

			if newProb > maxProb[nextNode] {
				maxProb[nextNode] = newProb
				heap.Push(pq, &Item{node: nextNode, prob: newProb})
			}
		}
	}

	return maxProb[end_node]
}

type Item struct {
	node int
	prob float64
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].prob > pq[j].prob }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
