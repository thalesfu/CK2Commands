package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加沙GazaBarony struct {
	feud.BaseBarony
}

var BGaza加沙 feud.Barony = &加沙GazaBarony{}

func init() {
	f := BGaza加沙.(*加沙GazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaza",
		TitleName: "加沙",
		TitleCode: "b_gaza",
	}
}
