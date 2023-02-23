package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金蒂尔KimtirBarony struct {
	feud.BaseBarony
}

var BKimtir金蒂尔 feud.Barony = &金蒂尔KimtirBarony{}

func init() {
	f := BKimtir金蒂尔.(*金蒂尔KimtirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kimtir",
		TitleName: "金蒂尔",
		TitleCode: "b_kimtir",
	}
}
