package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马扎卡MazacaBarony struct {
	feud.BaseBarony
}

var BMazaca马扎卡 feud.Barony = &马扎卡MazacaBarony{}

func init() {
	f := BMazaca马扎卡.(*马扎卡MazacaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazaca",
		TitleName: "马扎卡",
		TitleCode: "b_mazaca",
	}
}
