package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓罗德DunrodBarony struct {
	feud.BaseBarony
}

var BDunrod邓罗德 feud.Barony = &邓罗德DunrodBarony{}

func init() {
	f := BDunrod邓罗德.(*邓罗德DunrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunrod",
		TitleName: "邓罗德",
		TitleCode: "b_dunrod",
	}
}
