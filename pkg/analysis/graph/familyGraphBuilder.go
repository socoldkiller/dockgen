package analysis

import "dockgen/pkg/analysis/ir"

type FamilyGraphBuilder struct{}

func NewFamilyGraphBuilder() *FamilyGraphBuilder {
	return &FamilyGraphBuilder{}
}

func (b *FamilyGraphBuilder) BuildEdges(graph IRGraph) {
	nodes := graph.CmdList()

	familyGroups := make(map[string][]*ir.BashCommandIR)

	for _, node := range nodes {
		familyGroups[node.Program] = append(familyGroups[node.Program], node)
	}

	for _, group := range familyGroups {
		for i := 0; i < len(group)-1; i++ {
			from := group[i]
			to := group[i+1]
			BaseEdge{}.Add(graph, from.ID, to.ID, []string{"family"})
		}
	}
}
