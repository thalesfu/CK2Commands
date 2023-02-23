package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚多DiadoBarony struct {
	feud.BaseBarony
}

var BDiado迪亚多 feud.Barony = &迪亚多DiadoBarony{}

func init() {
	f := BDiado迪亚多.(*迪亚多DiadoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diado",
		TitleName: "迪亚多",
		TitleCode: "b_diado",
	}
}
