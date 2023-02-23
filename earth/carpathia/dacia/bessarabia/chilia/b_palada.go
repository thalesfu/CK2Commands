package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉德PaladaBarony struct {
	feud.BaseBarony
}

var BPalada帕拉德 feud.Barony = &帕拉德PaladaBarony{}

func init() {
	f := BPalada帕拉德.(*帕拉德PaladaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palada",
		TitleName: "帕拉德",
		TitleCode: "b_palada",
	}
}
