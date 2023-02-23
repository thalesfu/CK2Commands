package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里斯塔诺OristanoBarony struct {
	feud.BaseBarony
}

var BOristano奥里斯塔诺 feud.Barony = &奥里斯塔诺OristanoBarony{}

func init() {
	f := BOristano奥里斯塔诺.(*奥里斯塔诺OristanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oristano",
		TitleName: "奥里斯塔诺",
		TitleCode: "b_oristano",
	}
}
