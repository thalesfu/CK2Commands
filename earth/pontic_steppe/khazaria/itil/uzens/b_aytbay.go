package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾特拜AytbayBarony struct {
	feud.BaseBarony
}

var BAytbay艾特拜 feud.Barony = &艾特拜AytbayBarony{}

func init() {
    f := BAytbay艾特拜.(*艾特拜AytbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aytbay",
		TitleName: "艾特拜",
		TitleCode: "b_aytbay",
	}
}
