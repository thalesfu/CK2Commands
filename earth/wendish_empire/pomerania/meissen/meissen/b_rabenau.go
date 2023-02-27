package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉伯瑙RabenauBarony struct {
	feud.BaseBarony
}

var BRabenau拉伯瑙 feud.Barony = &拉伯瑙RabenauBarony{}

func init() {
    f := BRabenau拉伯瑙.(*拉伯瑙RabenauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabenau",
		TitleName: "拉伯瑙",
		TitleCode: "b_rabenau",
	}
}
