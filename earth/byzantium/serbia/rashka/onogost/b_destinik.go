package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德斯蒂尼克DestinikBarony struct {
	feud.BaseBarony
}

var BDestinik德斯蒂尼克 feud.Barony = &德斯蒂尼克DestinikBarony{}

func init() {
	f := BDestinik德斯蒂尼克.(*德斯蒂尼克DestinikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "destinik",
		TitleName: "德斯蒂尼克",
		TitleCode: "b_destinik",
	}
}
