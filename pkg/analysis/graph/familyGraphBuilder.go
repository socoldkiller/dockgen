package analysis

import "dockgen/pkg/analysis"

type FamilyGraphBuilder struct{}

func NewFamilyGraphBuilder() *FamilyGraphBuilder {
	return &FamilyGraphBuilder{}
}

func (b *FamilyGraphBuilder) BuildEdges(graph IRGraph) {
	nodes := graph.CmdList()

	familyGroups := make(map[string][]*analysis.CommandIR)

	for _, node := range nodes {
		familyGroups[node.Raw] = append(familyGroups[node.Raw], node)
	}

	for _, group := range familyGroups {
		for i := 0; i < len(group)-1; i++ {
			from := group[i]
			to := group[i+1]
			BaseEdge{}.Add(graph, from.ID, to.ID, []string{"family"})
		}
	}
}
