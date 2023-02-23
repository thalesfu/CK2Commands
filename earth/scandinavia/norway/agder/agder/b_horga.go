package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔加HorgaBarony struct {
	feud.BaseBarony
}

var BHorga霍尔加 feud.Barony = &霍尔加HorgaBarony{}

func init() {
	f := BHorga霍尔加.(*霍尔加HorgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horga",
		TitleName: "霍尔加",
		TitleCode: "b_horga",
	}
}
