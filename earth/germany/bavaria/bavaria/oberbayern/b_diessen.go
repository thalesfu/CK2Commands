package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪森DiessenBarony struct {
	feud.BaseBarony
}

var BDiessen迪森 feud.Barony = &迪森DiessenBarony{}

func init() {
	f := BDiessen迪森.(*迪森DiessenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diessen",
		TitleName: "迪森",
		TitleCode: "b_diessen",
	}
}
