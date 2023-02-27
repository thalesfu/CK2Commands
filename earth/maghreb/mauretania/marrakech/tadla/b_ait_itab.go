package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾特伊塔布Ait_itabBarony struct {
	feud.BaseBarony
}

var BAit_itab艾特伊塔布 feud.Barony = &艾特伊塔布Ait_itabBarony{}

func init() {
    f := BAit_itab艾特伊塔布.(*艾特伊塔布Ait_itabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ait_itab",
		TitleName: "艾特伊塔布",
		TitleCode: "b_ait_itab",
	}
}
