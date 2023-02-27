package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特勒斯ChartreuseBarony struct {
	feud.BaseBarony
}

var BChartreuse沙特勒斯 feud.Barony = &沙特勒斯ChartreuseBarony{}

func init() {
    f := BChartreuse沙特勒斯.(*沙特勒斯ChartreuseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chartreuse",
		TitleName: "沙特勒斯",
		TitleCode: "b_chartreuse",
	}
}
