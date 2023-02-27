package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔布赫摩斯KatabathmosBarony struct {
	feud.BaseBarony
}

var BKatabathmos卡塔布赫摩斯 feud.Barony = &卡塔布赫摩斯KatabathmosBarony{}

func init() {
    f := BKatabathmos卡塔布赫摩斯.(*卡塔布赫摩斯KatabathmosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katabathmos",
		TitleName: "卡塔布赫摩斯",
		TitleCode: "b_katabathmos",
	}
}
