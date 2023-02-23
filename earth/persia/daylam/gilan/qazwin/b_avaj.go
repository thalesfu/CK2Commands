package qazwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦吉AvajBarony struct {
	feud.BaseBarony
}

var BAvaj阿瓦吉 feud.Barony = &阿瓦吉AvajBarony{}

func init() {
	f := BAvaj阿瓦吉.(*阿瓦吉AvajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avaj",
		TitleName: "阿瓦吉",
		TitleCode: "b_avaj",
	}
}
