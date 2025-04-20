package analysis

func (b BaseGraphBuilder) BuildEdges(graph IRGraph) {
	for idx := range len(graph.CmdList()) - 1 {
		nowID := graph.CmdList()[idx].ID
		nextID := graph.CmdList()[idx+1].ID
		BaseEdge{}.Add(graph, nowID, nextID, []string{"context"})
	}
}
