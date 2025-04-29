package analysis

import "dockgen/pkg/analysis/types"

type BaseGraphBuilder struct {
}

func (b BaseGraphBuilder) BuildEdges(graph types.CFGraph) {
	for idx := range len(graph.CmdList()) - 1 {
		nowID := graph.CmdList()[idx].ID()
		nextID := graph.CmdList()[idx+1].ID()
		e := types.BashEdge{}
		e.Add(graph, nowID, nextID, []string{"context"})
	}

}
