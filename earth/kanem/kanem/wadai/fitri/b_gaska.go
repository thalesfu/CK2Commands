package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加斯卡GaskaBarony struct {
	feud.BaseBarony
}

var BGaska加斯卡 feud.Barony = &加斯卡GaskaBarony{}

func init() {
	f := BGaska加斯卡.(*加斯卡GaskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaska",
		TitleName: "加斯卡",
		TitleCode: "b_gaska",
	}
}
