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

func (g *BashGraph) Nodes() map[int]IR {
	return g.nodes
}

func (g *BashGraph) Edges() map[int][]Edge {
	return g.edges
}

func (g *BashGraph) CmdList() []IR {
	return g.irs
}

func (g *BashGraph) GetCmdGroup(program string, tag string) []IR {
	nodes := g.CmdList()
	var groups []IR
	idx := slices.IndexFunc(nodes, func(ir IR) bool { return ir.Program() == program })
	if idx == -1 {
		return nil
	}
	return g.getTagGroup(nodes[idx].ID(), tag, groups)
}

func (g *BashGraph) getTagGroup(from int, tag string, groups []IR) []IR {
	nowNode := g.Nodes()[from]
	if nowNode != nil {
		groups = append(groups, nowNode)
	}

	for _, edge := range g.edges[from] {
		nextID, _ := edge.To()
		if slices.Contains(edge.Tags(), tag) {
			groups = g.getTagGroup(nextID, tag, groups)
		}
	}

	return groups
}
