package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍尔奇BarcsBarony struct {
	feud.BaseBarony
}

var BBarcs鲍尔奇 feud.Barony = &鲍尔奇BarcsBarony{}

func init() {
    f := BBarcs鲍尔奇.(*鲍尔奇BarcsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barcs",
		TitleName: "鲍尔奇",
		TitleCode: "b_barcs",
	}
}
