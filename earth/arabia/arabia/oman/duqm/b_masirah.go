package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马西拉MasirahBarony struct {
	feud.BaseBarony
}

var BMasirah马西拉 feud.Barony = &马西拉MasirahBarony{}

func init() {
	f := BMasirah马西拉.(*马西拉MasirahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masirah",
		TitleName: "马西拉",
		TitleCode: "b_masirah",
	}
}
