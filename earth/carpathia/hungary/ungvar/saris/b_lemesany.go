package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱梅沙尼LemesanyBarony struct {
	feud.BaseBarony
}

var BLemesany莱梅沙尼 feud.Barony = &莱梅沙尼LemesanyBarony{}

func init() {
	f := BLemesany莱梅沙尼.(*莱梅沙尼LemesanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lemesany",
		TitleName: "莱梅沙尼",
		TitleCode: "b_lemesany",
	}
}
