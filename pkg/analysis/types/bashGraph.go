package types

import "slices"

type BashGraph struct {
	irs   []IR
	nodes map[int]IR
	edges map[int][]Edge
}

func NewBashGraph(irTokens []IR) *BashGraph {
	g := &BashGraph{
		irs:   irTokens,
		nodes: make(map[int]IR),
		edges: make(map[int][]Edge),
	}
	for _, ir := range irTokens {
		g.nodes[ir.ID()] = ir
	}
	return g
}

func (b *BashGraph) InitGraph(builders []GraphBuilder) {

	for _, builder := range builders {
		//todo
		builder.BuildEdges(b)
	}

}

func (b *BashGraph) Nodes() map[int]IR {
	return b.nodes
}

func (b *BashGraph) Edges() map[int][]Edge {
	return b.edges
}

func (b *BashGraph) CmdList() []IR {
	return b.irs
}

func (g *BashGraph) GetTagGroup(from int, tag string, groups []IR) []IR {
	nowNode := g.Nodes()[from]
	if nowNode != nil {
		groups = append(groups, nowNode)
	}

	for _, edge := range g.edges[from] {
		nextID, _ := edge.To()
		if slices.Contains(edge.Tags(), tag) {
			groups = g.GetTagGroup(nextID, tag, groups)
		}
	}

	return groups
}
