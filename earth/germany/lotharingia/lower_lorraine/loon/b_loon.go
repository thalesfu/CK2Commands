package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛翁LoonBarony struct {
	feud.BaseBarony
}

var BLoon洛翁 feud.Barony = &洛翁LoonBarony{}

func init() {
    f := BLoon洛翁.(*洛翁LoonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loon",
		TitleName: "洛翁",
		TitleCode: "b_loon",
	}
}
