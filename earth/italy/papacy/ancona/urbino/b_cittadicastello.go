package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰洛城CittadicastelloBarony struct {
	feud.BaseBarony
}

var BCittadicastello卡斯泰洛城 feud.Barony = &卡斯泰洛城CittadicastelloBarony{}

func init() {
    f := BCittadicastello卡斯泰洛城.(*卡斯泰洛城CittadicastelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cittadicastello",
		TitleName: "卡斯泰洛城",
		TitleCode: "b_cittadicastello",
	}
}
