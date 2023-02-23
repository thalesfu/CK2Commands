package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卑多斯AbydosBarony struct {
	feud.BaseBarony
}

var BAbydos阿卑多斯 feud.Barony = &阿卑多斯AbydosBarony{}

func init() {
	f := BAbydos阿卑多斯.(*阿卑多斯AbydosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abydos",
		TitleName: "阿卑多斯",
		TitleCode: "b_abydos",
	}
}
