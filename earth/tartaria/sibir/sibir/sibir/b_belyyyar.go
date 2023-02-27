package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白亚尔BelyyyarBarony struct {
	feud.BaseBarony
}

var BBelyyyar白亚尔 feud.Barony = &白亚尔BelyyyarBarony{}

func init() {
    f := BBelyyyar白亚尔.(*白亚尔BelyyyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belyyyar",
		TitleName: "白亚尔",
		TitleCode: "b_belyyyar",
	}
}
