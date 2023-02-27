package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克宁KninBarony struct {
	feud.BaseBarony
}

var BKnin克宁 feud.Barony = &克宁KninBarony{}

func init() {
    f := BKnin克宁.(*克宁KninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knin",
		TitleName: "克宁",
		TitleCode: "b_knin",
	}
}
