package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里尔LilleBarony struct {
	feud.BaseBarony
}

var BLille里尔 feud.Barony = &里尔LilleBarony{}

func init() {
    f := BLille里尔.(*里尔LilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lille",
		TitleName: "里尔",
		TitleCode: "b_lille",
	}
}
