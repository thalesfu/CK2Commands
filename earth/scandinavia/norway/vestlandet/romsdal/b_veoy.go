package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维伊VeoyBarony struct {
	feud.BaseBarony
}

var BVeoy维伊 feud.Barony = &维伊VeoyBarony{}

func init() {
	f := BVeoy维伊.(*维伊VeoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veoy",
		TitleName: "维伊",
		TitleCode: "b_veoy",
	}
}
