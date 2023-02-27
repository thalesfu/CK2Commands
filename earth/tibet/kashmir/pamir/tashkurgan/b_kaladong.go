package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉懂KaladongBarony struct {
	feud.BaseBarony
}

var BKaladong卡拉懂 feud.Barony = &卡拉懂KaladongBarony{}

func init() {
    f := BKaladong卡拉懂.(*卡拉懂KaladongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaladong",
		TitleName: "卡拉懂",
		TitleCode: "b_kaladong",
	}
}
