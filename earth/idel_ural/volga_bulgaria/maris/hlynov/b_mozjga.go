package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫日加MozjgaBarony struct {
	feud.BaseBarony
}

var BMozjga莫日加 feud.Barony = &莫日加MozjgaBarony{}

func init() {
    f := BMozjga莫日加.(*莫日加MozjgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mozjga",
		TitleName: "莫日加",
		TitleCode: "b_mozjga",
	}
}
