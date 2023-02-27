package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅沙瓦NieszawaBarony struct {
	feud.BaseBarony
}

var BNieszawa涅沙瓦 feud.Barony = &涅沙瓦NieszawaBarony{}

func init() {
    f := BNieszawa涅沙瓦.(*涅沙瓦NieszawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nieszawa",
		TitleName: "涅沙瓦",
		TitleCode: "b_nieszawa",
	}
}
