package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰加扎TaghazaBarony struct {
	feud.BaseBarony
}

var BTaghaza泰加扎 feud.Barony = &泰加扎TaghazaBarony{}

func init() {
	f := BTaghaza泰加扎.(*泰加扎TaghazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taghaza",
		TitleName: "泰加扎",
		TitleCode: "b_taghaza",
	}
}
