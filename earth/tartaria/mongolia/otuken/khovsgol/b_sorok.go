package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索罗克SorokBarony struct {
	feud.BaseBarony
}

var BSorok索罗克 feud.Barony = &索罗克SorokBarony{}

func init() {
    f := BSorok索罗克.(*索罗克SorokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorok",
		TitleName: "索罗克",
		TitleCode: "b_sorok",
	}
}
