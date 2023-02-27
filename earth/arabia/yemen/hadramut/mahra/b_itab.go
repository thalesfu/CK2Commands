package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊塔卜ItabBarony struct {
	feud.BaseBarony
}

var BItab伊塔卜 feud.Barony = &伊塔卜ItabBarony{}

func init() {
    f := BItab伊塔卜.(*伊塔卜ItabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itab",
		TitleName: "伊塔卜",
		TitleCode: "b_itab",
	}
}
