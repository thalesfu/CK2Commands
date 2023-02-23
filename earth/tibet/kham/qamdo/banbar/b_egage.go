package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俄嘎格EgageBarony struct {
	feud.BaseBarony
}

var BEgage俄嘎格 feud.Barony = &俄嘎格EgageBarony{}

func init() {
	f := BEgage俄嘎格.(*俄嘎格EgageBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egage",
		TitleName: "俄嘎格",
		TitleCode: "b_egage",
	}
}
