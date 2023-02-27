package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓巴DunbarBarony struct {
	feud.BaseBarony
}

var BDunbar邓巴 feud.Barony = &邓巴DunbarBarony{}

func init() {
    f := BDunbar邓巴.(*邓巴DunbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunbar",
		TitleName: "邓巴",
		TitleCode: "b_dunbar",
	}
}
