package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰加马赫LlangammarchBarony struct {
	feud.BaseBarony
}

var BLlangammarch兰加马赫 feud.Barony = &兰加马赫LlangammarchBarony{}

func init() {
	f := BLlangammarch兰加马赫.(*兰加马赫LlangammarchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llangammarch",
		TitleName: "兰加马赫",
		TitleCode: "b_llangammarch",
	}
}
