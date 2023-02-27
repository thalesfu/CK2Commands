package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓利DunollieBarony struct {
	feud.BaseBarony
}

var BDunollie邓利 feud.Barony = &邓利DunollieBarony{}

func init() {
    f := BDunollie邓利.(*邓利DunollieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunollie",
		TitleName: "邓利",
		TitleCode: "b_dunollie",
	}
}
