package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗里安KoloriangBarony struct {
	feud.BaseBarony
}

var BKoloriang克罗里安 feud.Barony = &克罗里安KoloriangBarony{}

func init() {
    f := BKoloriang克罗里安.(*克罗里安KoloriangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koloriang",
		TitleName: "克罗里安",
		TitleCode: "b_koloriang",
	}
}
