package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米勒福特MihrleftBarony struct {
	feud.BaseBarony
}

var BMihrleft米勒福特 feud.Barony = &米勒福特MihrleftBarony{}

func init() {
    f := BMihrleft米勒福特.(*米勒福特MihrleftBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mihrleft",
		TitleName: "米勒福特",
		TitleCode: "b_mihrleft",
	}
}
