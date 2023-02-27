package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑祇梨SanjeliBarony struct {
	feud.BaseBarony
}

var BSanjeli桑祇梨 feud.Barony = &桑祇梨SanjeliBarony{}

func init() {
    f := BSanjeli桑祇梨.(*桑祇梨SanjeliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanjeli",
		TitleName: "桑祇梨",
		TitleCode: "b_sanjeli",
	}
}
