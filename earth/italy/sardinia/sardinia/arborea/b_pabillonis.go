package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕比洛尼斯PabillonisBarony struct {
	feud.BaseBarony
}

var BPabillonis帕比洛尼斯 feud.Barony = &帕比洛尼斯PabillonisBarony{}

func init() {
	f := BPabillonis帕比洛尼斯.(*帕比洛尼斯PabillonisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pabillonis",
		TitleName: "帕比洛尼斯",
		TitleCode: "b_pabillonis",
	}
}
