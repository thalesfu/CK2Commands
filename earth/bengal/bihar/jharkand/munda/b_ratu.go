package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗兜RatuBarony struct {
	feud.BaseBarony
}

var BRatu罗兜 feud.Barony = &罗兜RatuBarony{}

func init() {
	f := BRatu罗兜.(*罗兜RatuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ratu",
		TitleName: "罗兜",
		TitleCode: "b_ratu",
	}
}
