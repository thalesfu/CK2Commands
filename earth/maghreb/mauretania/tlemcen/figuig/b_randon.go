package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗东RandonBarony struct {
	feud.BaseBarony
}

var BRandon朗东 feud.Barony = &朗东RandonBarony{}

func init() {
    f := BRandon朗东.(*朗东RandonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "randon",
		TitleName: "朗东",
		TitleCode: "b_randon",
	}
}
