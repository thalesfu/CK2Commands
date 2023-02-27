package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里瓦PrivasBarony struct {
	feud.BaseBarony
}

var BPrivas普里瓦 feud.Barony = &普里瓦PrivasBarony{}

func init() {
    f := BPrivas普里瓦.(*普里瓦PrivasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "privas",
		TitleName: "普里瓦",
		TitleCode: "b_privas",
	}
}
