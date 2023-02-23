package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希农堡ChateauchinonBarony struct {
	feud.BaseBarony
}

var BChateauchinon希农堡 feud.Barony = &希农堡ChateauchinonBarony{}

func init() {
	f := BChateauchinon希农堡.(*希农堡ChateauchinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateauchinon",
		TitleName: "希农堡",
		TitleCode: "b_chateauchinon",
	}
}
