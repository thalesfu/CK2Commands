package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉雷多LaredoBarony struct {
	feud.BaseBarony
}

var BLaredo拉雷多 feud.Barony = &拉雷多LaredoBarony{}

func init() {
    f := BLaredo拉雷多.(*拉雷多LaredoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laredo",
		TitleName: "拉雷多",
		TitleCode: "b_laredo",
	}
}
