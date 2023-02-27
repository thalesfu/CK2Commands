package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基克努尔KiknurBarony struct {
	feud.BaseBarony
}

var BKiknur基克努尔 feud.Barony = &基克努尔KiknurBarony{}

func init() {
    f := BKiknur基克努尔.(*基克努尔KiknurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiknur",
		TitleName: "基克努尔",
		TitleCode: "b_kiknur",
	}
}
