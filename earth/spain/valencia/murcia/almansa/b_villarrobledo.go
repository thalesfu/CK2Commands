package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚罗夫莱多VillarrobledoBarony struct {
	feud.BaseBarony
}

var BVillarrobledo比利亚罗夫莱多 feud.Barony = &比利亚罗夫莱多VillarrobledoBarony{}

func init() {
	f := BVillarrobledo比利亚罗夫莱多.(*比利亚罗夫莱多VillarrobledoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villarrobledo",
		TitleName: "比利亚罗夫莱多",
		TitleCode: "b_villarrobledo",
	}
}
