package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷布诺耶RybnoyeBarony struct {
	feud.BaseBarony
}

var BRybnoye雷布诺耶 feud.Barony = &雷布诺耶RybnoyeBarony{}

func init() {
	f := BRybnoye雷布诺耶.(*雷布诺耶RybnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rybnoye",
		TitleName: "雷布诺耶",
		TitleCode: "b_rybnoye",
	}
}
