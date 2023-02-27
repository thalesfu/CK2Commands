package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀拉库顺Kara_koshunBarony struct {
	feud.BaseBarony
}

var BKara_koshun喀拉库顺 feud.Barony = &喀拉库顺Kara_koshunBarony{}

func init() {
    f := BKara_koshun喀拉库顺.(*喀拉库顺Kara_koshunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_koshun",
		TitleName: "喀拉库顺",
		TitleCode: "b_kara_koshun",
	}
}
