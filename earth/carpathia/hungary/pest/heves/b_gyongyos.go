package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珍珠市GyongyosBarony struct {
	feud.BaseBarony
}

var BGyongyos珍珠市 feud.Barony = &珍珠市GyongyosBarony{}

func init() {
	f := BGyongyos珍珠市.(*珍珠市GyongyosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyongyos",
		TitleName: "珍珠市",
		TitleCode: "b_gyongyos",
	}
}
