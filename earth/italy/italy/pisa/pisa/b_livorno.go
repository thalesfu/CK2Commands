package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利沃诺LivornoBarony struct {
	feud.BaseBarony
}

var BLivorno利沃诺 feud.Barony = &利沃诺LivornoBarony{}

func init() {
	f := BLivorno利沃诺.(*利沃诺LivornoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "livorno",
		TitleName: "利沃诺",
		TitleCode: "b_livorno",
	}
}
