package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗莱什蒂FalestiBarony struct {
	feud.BaseBarony
}

var BFalesti弗莱什蒂 feud.Barony = &弗莱什蒂FalestiBarony{}

func init() {
	f := BFalesti弗莱什蒂.(*弗莱什蒂FalestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falesti",
		TitleName: "弗莱什蒂",
		TitleCode: "b_falesti",
	}
}
