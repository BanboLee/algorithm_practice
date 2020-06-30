package graph

/*
并查集实现
*/

type UnionFindSet struct {
	father []int
	rank   []int
}

func (ufs *UnionFindSet) Find(x int) int {
	if x != ufs.father[x] {
		ufs.father[x] = ufs.Find(ufs.father[x])
	}
	return ufs.father[x]
}

func (ufs *UnionFindSet) Union(x, y int) {
	x, y = ufs.Find(x), ufs.Find(y)
	if ufs.rank[x] > ufs.rank[y] {
		ufs.father[y] = x
	} else {
		ufs.father[x] = y
		if ufs.rank[x] == ufs.rank[y] {
			ufs.rank[x]++
		}
	}
}
