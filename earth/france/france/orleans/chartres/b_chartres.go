package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特尔ChartresBarony struct {
	feud.BaseBarony
}

var BChartres沙特尔 feud.Barony = &沙特尔ChartresBarony{}

func init() {
	f := BChartres沙特尔.(*沙特尔ChartresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chartres",
		TitleName: "沙特尔",
		TitleCode: "b_chartres",
	}
}
