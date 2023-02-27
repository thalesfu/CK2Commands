package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔蒂VertusBarony struct {
	feud.BaseBarony
}

var BVertus韦尔蒂 feud.Barony = &韦尔蒂VertusBarony{}

func init() {
    f := BVertus韦尔蒂.(*韦尔蒂VertusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vertus",
		TitleName: "韦尔蒂",
		TitleCode: "b_vertus",
	}
}
