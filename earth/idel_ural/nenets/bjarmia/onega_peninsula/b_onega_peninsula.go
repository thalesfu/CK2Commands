package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥涅加Onega_peninsulaBarony struct {
	feud.BaseBarony
}

var BOnega_peninsula奥涅加 feud.Barony = &奥涅加Onega_peninsulaBarony{}

func init() {
    f := BOnega_peninsula奥涅加.(*奥涅加Onega_peninsulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onega_peninsula",
		TitleName: "奥涅加",
		TitleCode: "b_onega_peninsula",
	}
}
