package types

type CFGraphBuilder interface {
	BuildEdges(graph CFGraph)
}

type CFGraph interface {
	InitGraph([]CFGraphBuilder)
	Nodes() map[int]IR
	Edges() map[int][]CFGEdge
	CmdList() []IR
}

type CFGEdge interface {
	Add(g CFGraph, from, to int, tag []string)
	To() (int, IR)
	From() (int, IR)
	Tags() []string
	SetTags(tags []string)
}
