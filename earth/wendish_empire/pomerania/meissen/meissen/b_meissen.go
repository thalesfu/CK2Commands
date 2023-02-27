package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈森MeissenBarony struct {
	feud.BaseBarony
}

var BMeissen迈森 feud.Barony = &迈森MeissenBarony{}

func init() {
    f := BMeissen迈森.(*迈森MeissenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meissen",
		TitleName: "迈森",
		TitleCode: "b_meissen",
	}
}
