package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约灵HjorringBarony struct {
	feud.BaseBarony
}

var BHjorring约灵 feud.Barony = &约灵HjorringBarony{}

func init() {
    f := BHjorring约灵.(*约灵HjorringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hjorring",
		TitleName: "约灵",
		TitleCode: "b_hjorring",
	}
}
