package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒孔特BarlecomteBarony struct {
	feud.BaseBarony
}

var BBarlecomte巴勒孔特 feud.Barony = &巴勒孔特BarlecomteBarony{}

func init() {
    f := BBarlecomte巴勒孔特.(*巴勒孔特BarlecomteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barlecomte",
		TitleName: "巴勒孔特",
		TitleCode: "b_barlecomte",
	}
}
