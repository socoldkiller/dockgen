package analysis

import (
	"dockgen/pkg/analysis/ir"
)

type GraphBuilder interface {
	BuildEdges(graph IRGraph)
}

type IRGraph interface {
	Init() error
	Build([]GraphBuilder) error
	Nodes() map[int]ir.IR
	Edges() map[int][]*BaseEdge
	CmdList() []ir.IR
}
type BaseGraphBuilder struct{}

type BaseIRGraph struct {
	irs      []ir.IR
	nodes    map[int]ir.IR
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
