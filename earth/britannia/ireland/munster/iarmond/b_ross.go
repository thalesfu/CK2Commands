package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯RossBarony struct {
	feud.BaseBarony
}

var BRoss罗斯 feud.Barony = &罗斯RossBarony{}

func init() {
	f := BRoss罗斯.(*罗斯RossBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ross",
		TitleName: "罗斯",
		TitleCode: "b_ross",
	}
}
