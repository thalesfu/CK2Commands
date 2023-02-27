package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 灵斯泰兹RingstedBarony struct {
	feud.BaseBarony
}

var BRingsted灵斯泰兹 feud.Barony = &灵斯泰兹RingstedBarony{}

func init() {
    f := BRingsted灵斯泰兹.(*灵斯泰兹RingstedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ringsted",
		TitleName: "灵斯泰兹",
		TitleCode: "b_ringsted",
	}
}
