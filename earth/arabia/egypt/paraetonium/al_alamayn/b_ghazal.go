package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加扎勒GhazalBarony struct {
	feud.BaseBarony
}

var BGhazal加扎勒 feud.Barony = &加扎勒GhazalBarony{}

func init() {
    f := BGhazal加扎勒.(*加扎勒GhazalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghazal",
		TitleName: "加扎勒",
		TitleCode: "b_ghazal",
	}
}
