package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞萨摩斯SesamusBarony struct {
	feud.BaseBarony
}

var BSesamus塞萨摩斯 feud.Barony = &塞萨摩斯SesamusBarony{}

func init() {
    f := BSesamus塞萨摩斯.(*塞萨摩斯SesamusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sesamus",
		TitleName: "塞萨摩斯",
		TitleCode: "b_sesamus",
	}
}
