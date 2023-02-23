package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于尔肯UlkenBarony struct {
	feud.BaseBarony
}

var BUlken于尔肯 feud.Barony = &于尔肯UlkenBarony{}

func init() {
	f := BUlken于尔肯.(*于尔肯UlkenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulken",
		TitleName: "于尔肯",
		TitleCode: "b_ulken",
	}
}
