package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆拘吒DevkotBarony struct {
	feud.BaseBarony
}

var BDevkot提婆拘吒 feud.Barony = &提婆拘吒DevkotBarony{}

func init() {
    f := BDevkot提婆拘吒.(*提婆拘吒DevkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devkot",
		TitleName: "提婆拘吒",
		TitleCode: "b_devkot",
	}
}
