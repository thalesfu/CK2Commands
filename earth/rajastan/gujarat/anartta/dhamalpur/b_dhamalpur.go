package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀摩罗补罗DhamalpurBarony struct {
	feud.BaseBarony
}

var BDhamalpur陀摩罗补罗 feud.Barony = &陀摩罗补罗DhamalpurBarony{}

func init() {
    f := BDhamalpur陀摩罗补罗.(*陀摩罗补罗DhamalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhamalpur",
		TitleName: "陀摩罗补罗",
		TitleCode: "b_dhamalpur",
	}
}
