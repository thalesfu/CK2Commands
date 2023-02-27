package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔讷夫VilleneuveBarony struct {
	feud.BaseBarony
}

var BVilleneuve维尔讷夫 feud.Barony = &维尔讷夫VilleneuveBarony{}

func init() {
    f := BVilleneuve维尔讷夫.(*维尔讷夫VilleneuveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villeneuve",
		TitleName: "维尔讷夫",
		TitleCode: "b_villeneuve",
	}
}
