package topological


type Graph struct {
	v int  // 顶点的个数
	linkedList [][]int // 邻接表
}

// s先于t，边s->t
func (g *Graph)addEdge(s int, t int)  {
	g.linkedList[s] = append(g.linkedList[s],t)
}

func newGraph(v int) *Graph{
	g := new(Graph)
	g.v = v
	g.linkedList = make([][]int,v)
	return g
}
