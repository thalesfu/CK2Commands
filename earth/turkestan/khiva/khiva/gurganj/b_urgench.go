package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉龙杰赤UrgenchBarony struct {
	feud.BaseBarony
}

var BUrgench玉龙杰赤 feud.Barony = &玉龙杰赤UrgenchBarony{}

func init() {
	f := BUrgench玉龙杰赤.(*玉龙杰赤UrgenchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urgench",
		TitleName: "玉龙杰赤",
		TitleCode: "b_urgench",
	}
}
