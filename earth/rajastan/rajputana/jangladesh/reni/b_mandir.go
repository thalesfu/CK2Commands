package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 寺庙MandirBarony struct {
	feud.BaseBarony
}

var BMandir寺庙 feud.Barony = &寺庙MandirBarony{}

func init() {
	f := BMandir寺庙.(*寺庙MandirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandir",
		TitleName: "寺庙",
		TitleCode: "b_mandir",
	}
}
