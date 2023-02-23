package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔特雷拉OltreraBarony struct {
	feud.BaseBarony
}

var BOltrera奥尔特雷拉 feud.Barony = &奥尔特雷拉OltreraBarony{}

func init() {
	f := BOltrera奥尔特雷拉.(*奥尔特雷拉OltreraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oltrera",
		TitleName: "奥尔特雷拉",
		TitleCode: "b_oltrera",
	}
}
