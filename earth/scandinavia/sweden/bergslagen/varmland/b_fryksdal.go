package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗吕克斯达尔FryksdalBarony struct {
	feud.BaseBarony
}

var BFryksdal弗吕克斯达尔 feud.Barony = &弗吕克斯达尔FryksdalBarony{}

func init() {
	f := BFryksdal弗吕克斯达尔.(*弗吕克斯达尔FryksdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fryksdal",
		TitleName: "弗吕克斯达尔",
		TitleCode: "b_fryksdal",
	}
}
