package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛伊吉什LoigisBarony struct {
	feud.BaseBarony
}

var BLoigis洛伊吉什 feud.Barony = &洛伊吉什LoigisBarony{}

func init() {
	f := BLoigis洛伊吉什.(*洛伊吉什LoigisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loigis",
		TitleName: "洛伊吉什",
		TitleCode: "b_loigis",
	}
}
