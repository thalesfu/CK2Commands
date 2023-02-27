package odessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列佐夫卡BerezivkaBarony struct {
	feud.BaseBarony
}

var BBerezivka别列佐夫卡 feud.Barony = &别列佐夫卡BerezivkaBarony{}

func init() {
    f := BBerezivka别列佐夫卡.(*别列佐夫卡BerezivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berezivka",
		TitleName: "别列佐夫卡",
		TitleCode: "b_berezivka",
	}
}
