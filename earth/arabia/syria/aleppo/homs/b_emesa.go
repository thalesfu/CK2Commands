package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃梅萨EmesaBarony struct {
	feud.BaseBarony
}

var BEmesa埃梅萨 feud.Barony = &埃梅萨EmesaBarony{}

func init() {
	f := BEmesa埃梅萨.(*埃梅萨EmesaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emesa",
		TitleName: "埃梅萨",
		TitleCode: "b_emesa",
	}
}
