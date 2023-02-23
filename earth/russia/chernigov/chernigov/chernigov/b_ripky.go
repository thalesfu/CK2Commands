package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里普基RipkyBarony struct {
	feud.BaseBarony
}

var BRipky里普基 feud.Barony = &里普基RipkyBarony{}

func init() {
	f := BRipky里普基.(*里普基RipkyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ripky",
		TitleName: "里普基",
		TitleCode: "b_ripky",
	}
}
