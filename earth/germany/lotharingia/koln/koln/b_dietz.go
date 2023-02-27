package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪茨DietzBarony struct {
	feud.BaseBarony
}

var BDietz迪茨 feud.Barony = &迪茨DietzBarony{}

func init() {
    f := BDietz迪茨.(*迪茨DietzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dietz",
		TitleName: "迪茨",
		TitleCode: "b_dietz",
	}
}
