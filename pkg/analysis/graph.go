package analysis

import "slices"

type GraphBuilder interface {
	BuildEdges(graph *IRGraph)
}

type BaseGraphBuilder struct {
}

func (b BaseGraphBuilder) BuildEdges(graph *IRGraph) {
	for idx := range len(graph.irs) - 1 {
		nowID := graph.irs[idx].ID
		nextID := graph.irs[idx+1].ID
		edge := AddIREdge(nowID, nextID, []string{"context"})
		graph.Edge[nowID] = append(graph.Edge[nowID], edge)
	}
}

type IRGraph struct {
	irs   []*CommandIR
	Nodes map[int]*CommandIR
	Edge  map[int][]*IREdge

	builders []GraphBuilder
}

func NewIRGraph(cmdList []*CommandIR, builders []GraphBuilder) *IRGraph {
	g := &IRGraph{
		irs:      cmdList,
		Nodes:    make(map[int]*CommandIR),
		Edge:     make(map[int][]*IREdge),
		builders: builders,
	}

	initNodes(g)
	initEdge(g)
	return g
}

func initNodes(graph *IRGraph) {
	for _, ir := range graph.irs {
		ID := ir.ID
		graph.Nodes[ID] = ir
	}
}

func initEdge(g *IRGraph) {
	for _, builder := range g.builders {
		builder.BuildEdges(g)
	}
}

func (g *IRGraph) GetTagGroup(ID int, tag string, groups []*CommandIR) []*CommandIR {
	for _, toIR := range g.Edge[ID] {
		toID := toIR.to
		if slices.Contains(toIR.Tags, tag) {
			groups = append(groups, g.Nodes[ID])
			return g.GetTagGroup(toID, tag, groups)
		}
	}
	return groups
}
