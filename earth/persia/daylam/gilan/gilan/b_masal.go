package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨拉MasalBarony struct {
	feud.BaseBarony
}

var BMasal马萨拉 feud.Barony = &马萨拉MasalBarony{}

func init() {
    f := BMasal马萨拉.(*马萨拉MasalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masal",
		TitleName: "马萨拉",
		TitleCode: "b_masal",
	}
}
