package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尼萨VenessaBarony struct {
	feud.BaseBarony
}

var BVenessa维尼萨 feud.Barony = &维尼萨VenessaBarony{}

func init() {
	f := BVenessa维尼萨.(*维尼萨VenessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "venessa",
		TitleName: "维尼萨",
		TitleCode: "b_venessa",
	}
}
