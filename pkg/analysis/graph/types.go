package analysis

import (
	"dockgen/pkg/analysis"
)

type GraphBuilder interface {
	BuildEdges(graph IRGraph)
}

type IRGraph interface {
	Init() error
	Build([]GraphBuilder) error
	Nodes() map[int]*analysis.CommandIR
	Edges() map[int][]*BaseEdge
	CmdList() []*analysis.CommandIR
}
type BaseGraphBuilder struct{}

type BaseIRGraph struct {
	irs      []*analysis.CommandIR
	nodes    map[int]*analysis.CommandIR
	edges    map[int][]*BaseEdge
	builders []GraphBuilder
}

type IREdge interface {
	Add(g IRGraph, from, to int, tag []string)
}

type BaseEdge struct {
	from   int
	to     int
	tags   []string
	weight int
}

func (b BaseEdge) Add(g IRGraph, from, to int, tag []string) {
	for _, next := range g.Edges()[from] {
		if next.to == to {
			next.tags = append(next.tags, tag...)
		}
	}
	g.Edges()[from] = append(g.Edges()[from], &BaseEdge{
		from: from,
		to:   to,
		tags: tag,
	})

}
