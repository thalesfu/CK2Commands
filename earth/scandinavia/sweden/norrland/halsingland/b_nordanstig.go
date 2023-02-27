package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努丹斯蒂格NordanstigBarony struct {
	feud.BaseBarony
}

var BNordanstig努丹斯蒂格 feud.Barony = &努丹斯蒂格NordanstigBarony{}

func init() {
    f := BNordanstig努丹斯蒂格.(*努丹斯蒂格NordanstigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordanstig",
		TitleName: "努丹斯蒂格",
		TitleCode: "b_nordanstig",
	}
}
