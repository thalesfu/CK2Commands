package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔戈沃耶KorgovoyeBarony struct {
	feud.BaseBarony
}

var BKorgovoye科尔戈沃耶 feud.Barony = &科尔戈沃耶KorgovoyeBarony{}

func init() {
    f := BKorgovoye科尔戈沃耶.(*科尔戈沃耶KorgovoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korgovoye",
		TitleName: "科尔戈沃耶",
		TitleCode: "b_korgovoye",
	}
}
