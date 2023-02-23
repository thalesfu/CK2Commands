package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚布利诺VillablinoBarony struct {
	feud.BaseBarony
}

var BVillablino比利亚布利诺 feud.Barony = &比利亚布利诺VillablinoBarony{}

func init() {
	f := BVillablino比利亚布利诺.(*比利亚布利诺VillablinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villablino",
		TitleName: "比利亚布利诺",
		TitleCode: "b_villablino",
	}
}
