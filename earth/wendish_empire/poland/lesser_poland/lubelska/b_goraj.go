package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈拉伊GorajBarony struct {
	feud.BaseBarony
}

var BGoraj戈拉伊 feud.Barony = &戈拉伊GorajBarony{}

func init() {
    f := BGoraj戈拉伊.(*戈拉伊GorajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goraj",
		TitleName: "戈拉伊",
		TitleCode: "b_goraj",
	}
}
