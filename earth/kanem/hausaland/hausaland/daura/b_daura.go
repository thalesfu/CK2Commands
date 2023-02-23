package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道拉DauraBarony struct {
	feud.BaseBarony
}

var BDaura道拉 feud.Barony = &道拉DauraBarony{}

func init() {
	f := BDaura道拉.(*道拉DauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daura",
		TitleName: "道拉",
		TitleCode: "b_daura",
	}
}
