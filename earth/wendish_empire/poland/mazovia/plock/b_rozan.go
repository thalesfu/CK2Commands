package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁然RozanBarony struct {
	feud.BaseBarony
}

var BRozan鲁然 feud.Barony = &鲁然RozanBarony{}

func init() {
    f := BRozan鲁然.(*鲁然RozanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozan",
		TitleName: "鲁然",
		TitleCode: "b_rozan",
	}
}
