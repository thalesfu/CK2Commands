package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉塔基亚LatakiahBarony struct {
	feud.BaseBarony
}

var BLatakiah拉塔基亚 feud.Barony = &拉塔基亚LatakiahBarony{}

func init() {
    f := BLatakiah拉塔基亚.(*拉塔基亚LatakiahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latakiah",
		TitleName: "拉塔基亚",
		TitleCode: "b_latakiah",
	}
}
