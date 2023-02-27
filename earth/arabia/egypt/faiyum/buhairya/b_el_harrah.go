package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉El_harrahBarony struct {
	feud.BaseBarony
}

var BEl_harrah哈拉 feud.Barony = &哈拉El_harrahBarony{}

func init() {
    f := BEl_harrah哈拉.(*哈拉El_harrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_harrah",
		TitleName: "哈拉",
		TitleCode: "b_el_harrah",
	}
}
