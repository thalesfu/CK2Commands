package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉善特RashaantBarony struct {
	feud.BaseBarony
}

var BRashaant拉善特 feud.Barony = &拉善特RashaantBarony{}

func init() {
    f := BRashaant拉善特.(*拉善特RashaantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rashaant",
		TitleName: "拉善特",
		TitleCode: "b_rashaant",
	}
}
