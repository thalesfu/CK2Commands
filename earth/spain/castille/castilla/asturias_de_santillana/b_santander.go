package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑坦德SantanderBarony struct {
	feud.BaseBarony
}

var BSantander桑坦德 feud.Barony = &桑坦德SantanderBarony{}

func init() {
    f := BSantander桑坦德.(*桑坦德SantanderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santander",
		TitleName: "桑坦德",
		TitleCode: "b_santander",
	}
}
