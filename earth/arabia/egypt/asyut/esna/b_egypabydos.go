package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拜多斯EgypabydosBarony struct {
	feud.BaseBarony
}

var BEgypabydos阿拜多斯 feud.Barony = &阿拜多斯EgypabydosBarony{}

func init() {
	f := BEgypabydos阿拜多斯.(*阿拜多斯EgypabydosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egypabydos",
		TitleName: "阿拜多斯",
		TitleCode: "b_egypabydos",
	}
}
