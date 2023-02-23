package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维克萨VyksaBarony struct {
	feud.BaseBarony
}

var BVyksa维克萨 feud.Barony = &维克萨VyksaBarony{}

func init() {
	f := BVyksa维克萨.(*维克萨VyksaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyksa",
		TitleName: "维克萨",
		TitleCode: "b_vyksa",
	}
}
