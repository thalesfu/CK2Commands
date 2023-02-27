package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉夫律塔KalavrytaBarony struct {
	feud.BaseBarony
}

var BKalavryta卡拉夫律塔 feud.Barony = &卡拉夫律塔KalavrytaBarony{}

func init() {
    f := BKalavryta卡拉夫律塔.(*卡拉夫律塔KalavrytaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalavryta",
		TitleName: "卡拉夫律塔",
		TitleCode: "b_kalavryta",
	}
}
