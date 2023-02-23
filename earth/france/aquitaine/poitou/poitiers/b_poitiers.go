package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普瓦捷PoitiersBarony struct {
	feud.BaseBarony
}

var BPoitiers普瓦捷 feud.Barony = &普瓦捷PoitiersBarony{}

func init() {
	f := BPoitiers普瓦捷.(*普瓦捷PoitiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poitiers",
		TitleName: "普瓦捷",
		TitleCode: "b_poitiers",
	}
}
