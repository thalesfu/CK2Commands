package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰戈伦LlangollenBarony struct {
	feud.BaseBarony
}

var BLlangollen兰戈伦 feud.Barony = &兰戈伦LlangollenBarony{}

func init() {
	f := BLlangollen兰戈伦.(*兰戈伦LlangollenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llangollen",
		TitleName: "兰戈伦",
		TitleCode: "b_llangollen",
	}
}
