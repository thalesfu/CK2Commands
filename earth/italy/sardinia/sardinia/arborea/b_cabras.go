package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡布拉斯CabrasBarony struct {
	feud.BaseBarony
}

var BCabras卡布拉斯 feud.Barony = &卡布拉斯CabrasBarony{}

func init() {
	f := BCabras卡布拉斯.(*卡布拉斯CabrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cabras",
		TitleName: "卡布拉斯",
		TitleCode: "b_cabras",
	}
}
