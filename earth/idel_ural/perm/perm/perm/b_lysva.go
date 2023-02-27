package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷西瓦LysvaBarony struct {
	feud.BaseBarony
}

var BLysva雷西瓦 feud.Barony = &雷西瓦LysvaBarony{}

func init() {
    f := BLysva雷西瓦.(*雷西瓦LysvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lysva",
		TitleName: "雷西瓦",
		TitleCode: "b_lysva",
	}
}
