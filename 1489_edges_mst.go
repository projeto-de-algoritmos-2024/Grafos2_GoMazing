import (
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	findMST := func(include, exclude int) int {
		uf := newUnionFind(n)
		totalWeight := 0
		if include != -1 {
			edge := edges[include]
			if uf.union(edge[0], edge[1]) {
				totalWeight += edge[2]
			}
		}
		for i, edge := range edges {
			if i == exclude {
				continue
			}
			if uf.union(edge[0], edge[1]) {
				totalWeight += edge[2]
			}
		}

		if uf.count() > 1 {
			return 1<<31 - 1
		}
		return totalWeight
	}

	originalWeight := findMST(-1, -1)

	critical := []int{}
	pseudoCritical := []int{}

	for i, edge := range edges {

		if findMST(-1, i) > originalWeight {
			critical = append(critical, edge[3])
		} else if findMST(i, -1) == originalWeight {

			pseudoCritical = append(pseudoCritical, edge[3])
		}
	}

	return [][]int{critical, pseudoCritical}
}

type unionFind struct {
	parent, rank []int
	count        int
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{parent, rank, size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return false
	}
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	uf.count--
	return true
}

func (uf *unionFind) count() int {
	return uf.count
}
