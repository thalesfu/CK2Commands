package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃斯坦VostanBarony struct {
	feud.BaseBarony
}

var BVostan沃斯坦 feud.Barony = &沃斯坦VostanBarony{}

func init() {
    f := BVostan沃斯坦.(*沃斯坦VostanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vostan",
		TitleName: "沃斯坦",
		TitleCode: "b_vostan",
	}
}
