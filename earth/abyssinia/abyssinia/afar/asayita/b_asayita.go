package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塞塔AsayitaBarony struct {
	feud.BaseBarony
}

var BAsayita阿塞塔 feud.Barony = &阿塞塔AsayitaBarony{}

func init() {
	f := BAsayita阿塞塔.(*阿塞塔AsayitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asayita",
		TitleName: "阿塞塔",
		TitleCode: "b_asayita",
	}
}
