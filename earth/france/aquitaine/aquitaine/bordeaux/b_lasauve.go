package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉索夫LasauveBarony struct {
	feud.BaseBarony
}

var BLasauve拉索夫 feud.Barony = &拉索夫LasauveBarony{}

func init() {
	f := BLasauve拉索夫.(*拉索夫LasauveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lasauve",
		TitleName: "拉索夫",
		TitleCode: "b_lasauve",
	}
}
