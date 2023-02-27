package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马蒂恩MotynBarony struct {
	feud.BaseBarony
}

var BMotyn马蒂恩 feud.Barony = &马蒂恩MotynBarony{}

func init() {
    f := BMotyn马蒂恩.(*马蒂恩MotynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motyn",
		TitleName: "马蒂恩",
		TitleCode: "b_motyn",
	}
}
