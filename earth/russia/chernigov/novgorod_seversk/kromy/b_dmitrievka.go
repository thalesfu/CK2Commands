package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季米特里耶夫卡DmitrievkaBarony struct {
	feud.BaseBarony
}

var BDmitrievka季米特里耶夫卡 feud.Barony = &季米特里耶夫卡DmitrievkaBarony{}

func init() {
    f := BDmitrievka季米特里耶夫卡.(*季米特里耶夫卡DmitrievkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dmitrievka",
		TitleName: "季米特里耶夫卡",
		TitleCode: "b_dmitrievka",
	}
}
