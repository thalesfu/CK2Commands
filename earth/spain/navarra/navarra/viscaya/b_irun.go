package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊伦IrunBarony struct {
	feud.BaseBarony
}

var BIrun伊伦 feud.Barony = &伊伦IrunBarony{}

func init() {
	f := BIrun伊伦.(*伊伦IrunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irun",
		TitleName: "伊伦",
		TitleCode: "b_irun",
	}
}
