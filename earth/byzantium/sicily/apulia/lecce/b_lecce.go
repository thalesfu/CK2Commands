package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱切LecceBarony struct {
	feud.BaseBarony
}

var BLecce莱切 feud.Barony = &莱切LecceBarony{}

func init() {
	f := BLecce莱切.(*莱切LecceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lecce",
		TitleName: "莱切",
		TitleCode: "b_lecce",
	}
}
