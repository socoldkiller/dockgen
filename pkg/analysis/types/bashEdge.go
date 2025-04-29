package types

type BashEdge struct {
	from   int
	to     int
	tags   []string
	weight int
	g      CFGraph
}

func (e *BashEdge) SetTags(tags []string) {
	e.tags = append(e.tags, tags...)
}

func (e *BashEdge) To() (int, IR) {
	return e.to, e.g.Nodes()[e.to]
}

func (e *BashEdge) From() (int, IR) {
	return e.from, e.g.Nodes()[e.from]
}

func (e *BashEdge) Tags() []string {
	return e.tags
}

func (e *BashEdge) Add(g CFGraph, from, to int, tag []string) {
	for _, edge := range g.Edges()[from] {
		toID, _ := edge.To()
		if toID == to {
			edge.SetTags(tag)
			return
		}
	}

	g.Edges()[from] = append(g.Edges()[from], &BashEdge{
		from: from,
		to:   to,
		tags: tag,
		g:    g,
	})

}
