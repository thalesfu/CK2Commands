package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰罗斯LlanrhosBarony struct {
	feud.BaseBarony
}

var BLlanrhos兰罗斯 feud.Barony = &兰罗斯LlanrhosBarony{}

func init() {
	f := BLlanrhos兰罗斯.(*兰罗斯LlanrhosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanrhos",
		TitleName: "兰罗斯",
		TitleCode: "b_llanrhos",
	}
}
