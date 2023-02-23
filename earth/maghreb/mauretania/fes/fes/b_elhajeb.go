package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈杰卜ElhajebBarony struct {
	feud.BaseBarony
}

var BElhajeb哈杰卜 feud.Barony = &哈杰卜ElhajebBarony{}

func init() {
	f := BElhajeb哈杰卜.(*哈杰卜ElhajebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elhajeb",
		TitleName: "哈杰卜",
		TitleCode: "b_elhajeb",
	}
}
