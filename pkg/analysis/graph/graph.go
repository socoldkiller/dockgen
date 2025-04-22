package analysis

import (
	"dockgen/pkg/analysis/ir"
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

func (g *BaseIRGraph) Nodes() map[int]*ir.BashCommandIR {
	return g.nodes
}

func (g *BaseIRGraph) Edges() map[int][]*BaseEdge {
	return g.edges
}

func (g *BaseIRGraph) CmdList() []*ir.BashCommandIR {
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

func NewIRGraph(cmdList []*ir.BashCommandIR) *BaseIRGraph {
	g := &BaseIRGraph{
		irs:   cmdList,
		nodes: make(map[int]*ir.BashCommandIR),
		edges: make(map[int][]*BaseEdge),
	}
	return g
}

func (g *BaseIRGraph) GetTagGroup(from int, tag string) []*ir.BashCommandIR {
	var (
		groups []*ir.BashCommandIR
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
