package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥梯罗MithilaBarony struct {
	feud.BaseBarony
}

var BMithila弥梯罗 feud.Barony = &弥梯罗MithilaBarony{}

func init() {
    f := BMithila弥梯罗.(*弥梯罗MithilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mithila",
		TitleName: "弥梯罗",
		TitleCode: "b_mithila",
	}
}
