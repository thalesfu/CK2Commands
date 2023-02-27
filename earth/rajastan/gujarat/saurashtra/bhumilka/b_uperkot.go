package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优波罗拘吒UperkotBarony struct {
	feud.BaseBarony
}

var BUperkot优波罗拘吒 feud.Barony = &优波罗拘吒UperkotBarony{}

func init() {
    f := BUperkot优波罗拘吒.(*优波罗拘吒UperkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uperkot",
		TitleName: "优波罗拘吒",
		TitleCode: "b_uperkot",
	}
}
