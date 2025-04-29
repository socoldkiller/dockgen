package types

type BashGraph struct {
	irs   []IR
	nodes map[int]IR
	edges map[int][]CFGEdge
}

func NewBashGraph(irTokens []IR) *BashGraph {
	g := &BashGraph{
		irs:   irTokens,
		nodes: make(map[int]IR),
		edges: make(map[int][]CFGEdge),
	}
	for _, ir := range irTokens {
		g.nodes[ir.ID()] = ir
	}
	return g
}

func (b *BashGraph) InitGraph(builders []CFGraphBuilder) {

	for _, builder := range builders {
		//todo
		builder.BuildEdges(b)
	}

}

func (g *BashGraph) Nodes() map[int]IR {
	return g.nodes
}

func (g *BashGraph) Edges() map[int][]CFGEdge {
	return g.edges
}

func (g *BashGraph) CmdList() []IR {
	return g.irs
}
