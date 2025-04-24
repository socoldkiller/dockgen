package types

type GraphBuilder interface {
	BuildEdges(graph IRGraph)
}

type IRGraph interface {
	InitGraph([]GraphBuilder)
	Nodes() map[int]IR
	Edges() map[int][]Edge
	CmdList() []IR
}

type Edge interface {
	Add(g IRGraph, from, to int, tag []string)
	To() (int, IR)
	From() (int, IR)
	Tags() []string
	SetTags(tags []string)
}
