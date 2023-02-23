package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯凡达克EsfandakBarony struct {
	feud.BaseBarony
}

var BEsfandak埃斯凡达克 feud.Barony = &埃斯凡达克EsfandakBarony{}

func init() {
	f := BEsfandak埃斯凡达克.(*埃斯凡达克EsfandakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esfandak",
		TitleName: "埃斯凡达克",
		TitleCode: "b_esfandak",
	}
}
