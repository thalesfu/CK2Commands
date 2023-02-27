package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考梅捷伊KaumeteyBarony struct {
	feud.BaseBarony
}

var BKaumetey考梅捷伊 feud.Barony = &考梅捷伊KaumeteyBarony{}

func init() {
    f := BKaumetey考梅捷伊.(*考梅捷伊KaumeteyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaumetey",
		TitleName: "考梅捷伊",
		TitleCode: "b_kaumetey",
	}
}
