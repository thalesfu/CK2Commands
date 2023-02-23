package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔扎TazaBarony struct {
	feud.BaseBarony
}

var BTaza塔扎 feud.Barony = &塔扎TazaBarony{}

func init() {
	f := BTaza塔扎.(*塔扎TazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taza",
		TitleName: "塔扎",
		TitleCode: "b_taza",
	}
}
