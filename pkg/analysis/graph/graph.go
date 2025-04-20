package analysis

import (
	"dockgen/pkg/analysis"
	"slices"
)

func initNodes(graph *BaseIRGraph) {
	for _, ir := range graph.irs {
		ID := ir.ID
		graph.nodes[ID] = ir
	}
}

func initEdge(g *BaseIRGraph) {
	for _, builder := range g.builders {
		builder.BuildEdges(g)
	}
}

func (g *BaseIRGraph) Nodes() map[int]*analysis.CommandIR {
	return g.nodes
}

func (g *BaseIRGraph) Edges() map[int][]*BaseEdge {
	return g.edges
}

func (g *BaseIRGraph) CmdList() []*analysis.CommandIR {
	return g.irs
}

func (g *BaseIRGraph) Init() error {
	initNodes(g)
	initEdge(g)
	return nil
}

func (g *BaseIRGraph) Build(builders []GraphBuilder) error {
	g.builders = builders
	return nil
}

func NewIRGraph(cmdList []*analysis.CommandIR) *BaseIRGraph {
	g := &BaseIRGraph{
		irs:   cmdList,
		nodes: make(map[int]*analysis.CommandIR),
		edges: make(map[int][]*BaseEdge),
	}
	return g
}

func (g *BaseIRGraph) GetTagGroup(from int, tag string) []*analysis.CommandIR {
	var (
		groups []*analysis.CommandIR
	)

	var dfs func(cur int)
	dfs = func(cur int) {
		if g.nodes[cur] == nil {
			return
		}
		groups = append(groups, g.nodes[cur])
		for _, next := range g.edges[cur] {
			if slices.Contains(next.tags, tag) {
				dfs(next.to)
			}
		}
	}
	dfs(from)
	return groups
}
