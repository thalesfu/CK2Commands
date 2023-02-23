package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃利奇ElcheBarony struct {
	feud.BaseBarony
}

var BElche埃利奇 feud.Barony = &埃利奇ElcheBarony{}

func init() {
	f := BElche埃利奇.(*埃利奇ElcheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elche",
		TitleName: "埃利奇",
		TitleCode: "b_elche",
	}
}
