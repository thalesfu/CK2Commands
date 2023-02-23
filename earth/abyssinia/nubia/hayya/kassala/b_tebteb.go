package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特卜特卜TebtebBarony struct {
	feud.BaseBarony
}

var BTebteb特卜特卜 feud.Barony = &特卜特卜TebtebBarony{}

func init() {
	f := BTebteb特卜特卜.(*特卜特卜TebtebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tebteb",
		TitleName: "特卜特卜",
		TitleCode: "b_tebteb",
	}
}
