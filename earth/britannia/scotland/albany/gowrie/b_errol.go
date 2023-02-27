package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃勒尔ErrolBarony struct {
	feud.BaseBarony
}

var BErrol埃勒尔 feud.Barony = &埃勒尔ErrolBarony{}

func init() {
    f := BErrol埃勒尔.(*埃勒尔ErrolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "errol",
		TitleName: "埃勒尔",
		TitleCode: "b_errol",
	}
}
