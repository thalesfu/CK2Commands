package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅夫拉NieblaBarony struct {
	feud.BaseBarony
}

var BNiebla涅夫拉 feud.Barony = &涅夫拉NieblaBarony{}

func init() {
	f := BNiebla涅夫拉.(*涅夫拉NieblaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niebla",
		TitleName: "涅夫拉",
		TitleCode: "b_niebla",
	}
}
