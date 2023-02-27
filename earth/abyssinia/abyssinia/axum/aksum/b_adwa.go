package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿杜瓦AdwaBarony struct {
	feud.BaseBarony
}

var BAdwa阿杜瓦 feud.Barony = &阿杜瓦AdwaBarony{}

func init() {
    f := BAdwa阿杜瓦.(*阿杜瓦AdwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adwa",
		TitleName: "阿杜瓦",
		TitleCode: "b_adwa",
	}
}
