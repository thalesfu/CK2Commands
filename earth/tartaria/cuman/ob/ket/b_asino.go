package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿西诺AsinoBarony struct {
	feud.BaseBarony
}

var BAsino阿西诺 feud.Barony = &阿西诺AsinoBarony{}

func init() {
    f := BAsino阿西诺.(*阿西诺AsinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asino",
		TitleName: "阿西诺",
		TitleCode: "b_asino",
	}
}
