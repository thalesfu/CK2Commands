package kyzyl_su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦扎AwazaBarony struct {
	feud.BaseBarony
}

var BAwaza阿瓦扎 feud.Barony = &阿瓦扎AwazaBarony{}

func init() {
    f := BAwaza阿瓦扎.(*阿瓦扎AwazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awaza",
		TitleName: "阿瓦扎",
		TitleCode: "b_awaza",
	}
}
