package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科龙KoronBarony struct {
	feud.BaseBarony
}

var BKoron科龙 feud.Barony = &科龙KoronBarony{}

func init() {
    f := BKoron科龙.(*科龙KoronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koron",
		TitleName: "科龙",
		TitleCode: "b_koron",
	}
}
