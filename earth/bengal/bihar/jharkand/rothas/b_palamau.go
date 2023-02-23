package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗摩郁PalamauBarony struct {
	feud.BaseBarony
}

var BPalamau波罗摩郁 feud.Barony = &波罗摩郁PalamauBarony{}

func init() {
	f := BPalamau波罗摩郁.(*波罗摩郁PalamauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palamau",
		TitleName: "波罗摩郁",
		TitleCode: "b_palamau",
	}
}
