package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克柳奇KljucBarony struct {
	feud.BaseBarony
}

var BKljuc克柳奇 feud.Barony = &克柳奇KljucBarony{}

func init() {
	f := BKljuc克柳奇.(*克柳奇KljucBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kljuc",
		TitleName: "克柳奇",
		TitleCode: "b_kljuc",
	}
}
