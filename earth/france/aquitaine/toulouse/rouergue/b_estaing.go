package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯坦EstaingBarony struct {
	feud.BaseBarony
}

var BEstaing埃斯坦 feud.Barony = &埃斯坦EstaingBarony{}

func init() {
	f := BEstaing埃斯坦.(*埃斯坦EstaingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estaing",
		TitleName: "埃斯坦",
		TitleCode: "b_estaing",
	}
}
