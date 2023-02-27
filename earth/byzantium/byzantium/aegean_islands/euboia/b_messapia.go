package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅萨皮亚MessapiaBarony struct {
	feud.BaseBarony
}

var BMessapia梅萨皮亚 feud.Barony = &梅萨皮亚MessapiaBarony{}

func init() {
    f := BMessapia梅萨皮亚.(*梅萨皮亚MessapiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "messapia",
		TitleName: "梅萨皮亚",
		TitleCode: "b_messapia",
	}
}
