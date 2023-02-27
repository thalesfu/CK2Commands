package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒瓦AlwaBarony struct {
	feud.BaseBarony
}

var BAlwa阿勒瓦 feud.Barony = &阿勒瓦AlwaBarony{}

func init() {
    f := BAlwa阿勒瓦.(*阿勒瓦AlwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alwa",
		TitleName: "阿勒瓦",
		TitleCode: "b_alwa",
	}
}
