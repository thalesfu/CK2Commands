package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓罗宾DunrobinBarony struct {
	feud.BaseBarony
}

var BDunrobin邓罗宾 feud.Barony = &邓罗宾DunrobinBarony{}

func init() {
	f := BDunrobin邓罗宾.(*邓罗宾DunrobinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunrobin",
		TitleName: "邓罗宾",
		TitleCode: "b_dunrobin",
	}
}
