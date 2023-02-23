package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈奥蒂亚NaeoteaBarony struct {
	feud.BaseBarony
}

var BNaeotea奈奥蒂亚 feud.Barony = &奈奥蒂亚NaeoteaBarony{}

func init() {
	f := BNaeotea奈奥蒂亚.(*奈奥蒂亚NaeoteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naeotea",
		TitleName: "奈奥蒂亚",
		TitleCode: "b_naeotea",
	}
}
