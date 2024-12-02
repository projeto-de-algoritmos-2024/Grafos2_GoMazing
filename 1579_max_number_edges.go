package main

func maxNumEdgesToRemove(n int, edges [][]int) int {

	type UnionFind struct {
		parent, rank []int
	}

	initUnionFind := func(size int) UnionFind {
		uf := UnionFind{
			parent: make([]int, size),
			rank:   make([]int, size),
		}
		for i := range uf.parent {
			uf.parent[i] = i
		}
		return uf
	}

	find := func(uf *UnionFind, x int) int {
		if uf.parent[x] != x {
			uf.parent[x] = find(uf, uf.parent[x])
		}
		return uf.parent[x]
	}

	union := func(uf *UnionFind, x, y int) bool {
		rootX := find(uf, x)
		rootY := find(uf, y)
		if rootX != rootY {
			if uf.rank[rootX] > uf.rank[rootY] {
				uf.parent[rootY] = rootX
			} else if uf.rank[rootX] < uf.rank[rootY] {
				uf.parent[rootX] = rootY
			} else {
				uf.parent[rootY] = rootX
				uf.rank[rootX]++
			}
			return true
		}
		return false
	}

	ufAlice := initUnionFind(n + 1)
	ufBob := initUnionFind(n + 1)

	edgesUsed := 0
	for _, edge := range edges {
		if edge[0] == 3 {
			if union(&ufAlice, edge[1], edge[2]) {
				union(&ufBob, edge[1], edge[2])
				edgesUsed++
			}
		}
	}

	for _, edge := range edges {
		if edge[0] == 1 {
			if union(&ufAlice, edge[1], edge[2]) {
				edgesUsed++
			}
		}
	}

	for _, edge := range edges {
		if edge[0] == 2 {
			if union(&ufBob, edge[1], edge[2]) {
				edgesUsed++
			}
		}
	}

	canTraverse := func(uf UnionFind) bool {
		root := find(&uf, 1)
		for i := 2; i <= n; i++ {
			if find(&uf, i) != root {
				return false
			}
		}
		return true
	}

	if canTraverse(ufAlice) && canTraverse(ufBob) {
		return len(edges) - edgesUsed
	}
	return -1
}
