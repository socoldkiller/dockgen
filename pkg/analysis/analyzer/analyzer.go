package analyzer

import (
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/log"
	"github.com/samber/lo"
)

type Result struct {
	name  string
	oldIR []types.IR
	newIR []types.IR

	qIR map[int]Item
}

func NewResult(name string, oldIR []types.IR, newIR []types.IR) *Result {
	if len(oldIR) != len(newIR) {
		log.Fatalf("The number of IRs must be equa")
	}
	qIR := make(map[int]Item)
	for idx, ir := range oldIR {
		qIR[ir.ID()] = Item{
			OldIR: oldIR[idx],
			NewIR: newIR[idx],
		}

	}
	return &Result{
		name:  name,
		oldIR: oldIR,
		newIR: newIR,
		qIR:   qIR,
	}

}

func (r Result) Name() string {
	return r.name
}

func (r Result) Raw() []string {
	return lo.Map(r.oldIR, func(raw types.IR, _ int) string {
		return raw.Cmd()
	})
}

func (r Result) Items() []Item {
	irs := lo.Zip2(r.oldIR, r.newIR)
	return lo.Map(irs, func(ir lo.Tuple2[types.IR, types.IR], index int) Item {
		return Item{
			OldIR: ir.A,
			NewIR: ir.B,
		}

	})
}

func (r Result) AnalyzeRaw() []string {
	return lo.Map(r.newIR, func(raw types.IR, _ int) string {
		return raw.Cmd()
	})
}

func (r Result) QueryIDIR(i int) (Item, bool) {
	ir, ok := r.qIR[i]
	return ir, ok
}
