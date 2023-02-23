package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟诃梨MuhriBarony struct {
	feud.BaseBarony
}

var BMuhri牟诃梨 feud.Barony = &牟诃梨MuhriBarony{}

func init() {
	f := BMuhri牟诃梨.(*牟诃梨MuhriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muhri",
		TitleName: "牟诃梨",
		TitleCode: "b_muhri",
	}
}
