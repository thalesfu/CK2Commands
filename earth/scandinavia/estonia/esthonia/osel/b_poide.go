package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀伊德PoideBarony struct {
	feud.BaseBarony
}

var BPoide珀伊德 feud.Barony = &珀伊德PoideBarony{}

func init() {
	f := BPoide珀伊德.(*珀伊德PoideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poide",
		TitleName: "珀伊德",
		TitleCode: "b_poide",
	}
}
