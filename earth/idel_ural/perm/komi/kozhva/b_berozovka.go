package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别罗佐夫卡BerozovkaBarony struct {
	feud.BaseBarony
}

var BBerozovka别罗佐夫卡 feud.Barony = &别罗佐夫卡BerozovkaBarony{}

func init() {
    f := BBerozovka别罗佐夫卡.(*别罗佐夫卡BerozovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berozovka",
		TitleName: "别罗佐夫卡",
		TitleCode: "b_berozovka",
	}
}
