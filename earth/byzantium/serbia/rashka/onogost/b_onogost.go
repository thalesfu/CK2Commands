package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥诺戈什特OnogostBarony struct {
	feud.BaseBarony
}

var BOnogost奥诺戈什特 feud.Barony = &奥诺戈什特OnogostBarony{}

func init() {
    f := BOnogost奥诺戈什特.(*奥诺戈什特OnogostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onogost",
		TitleName: "奥诺戈什特",
		TitleCode: "b_onogost",
	}
}
