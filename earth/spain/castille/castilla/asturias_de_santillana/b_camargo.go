package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马戈CamargoBarony struct {
	feud.BaseBarony
}

var BCamargo卡马戈 feud.Barony = &卡马戈CamargoBarony{}

func init() {
    f := BCamargo卡马戈.(*卡马戈CamargoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "camargo",
		TitleName: "卡马戈",
		TitleCode: "b_camargo",
	}
}
