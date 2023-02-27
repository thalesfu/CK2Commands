package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白垭BaiyaBarony struct {
	feud.BaseBarony
}

var BBaiya白垭 feud.Barony = &白垭BaiyaBarony{}

func init() {
    f := BBaiya白垭.(*白垭BaiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baiya",
		TitleName: "白垭",
		TitleCode: "b_baiya",
	}
}
