package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑威奇SandwichBarony struct {
	feud.BaseBarony
}

var BSandwich桑威奇 feud.Barony = &桑威奇SandwichBarony{}

func init() {
    f := BSandwich桑威奇.(*桑威奇SandwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandwich",
		TitleName: "桑威奇",
		TitleCode: "b_sandwich",
	}
}
