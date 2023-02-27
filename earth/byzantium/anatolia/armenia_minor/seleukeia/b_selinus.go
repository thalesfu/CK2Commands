package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞利努斯SelinusBarony struct {
	feud.BaseBarony
}

var BSelinus塞利努斯 feud.Barony = &塞利努斯SelinusBarony{}

func init() {
    f := BSelinus塞利努斯.(*塞利努斯SelinusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selinus",
		TitleName: "塞利努斯",
		TitleCode: "b_selinus",
	}
}
