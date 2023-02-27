package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺雷尼亚NorenaBarony struct {
	feud.BaseBarony
}

var BNorena诺雷尼亚 feud.Barony = &诺雷尼亚NorenaBarony{}

func init() {
    f := BNorena诺雷尼亚.(*诺雷尼亚NorenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norena",
		TitleName: "诺雷尼亚",
		TitleCode: "b_norena",
	}
}
