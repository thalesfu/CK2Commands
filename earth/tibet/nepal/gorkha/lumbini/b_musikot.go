package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟斯拘吒MusikotBarony struct {
	feud.BaseBarony
}

var BMusikot牟斯拘吒 feud.Barony = &牟斯拘吒MusikotBarony{}

func init() {
	f := BMusikot牟斯拘吒.(*牟斯拘吒MusikotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "musikot",
		TitleName: "牟斯拘吒",
		TitleCode: "b_musikot",
	}
}
