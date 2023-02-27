package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉约瓦CraiovaBarony struct {
	feud.BaseBarony
}

var BCraiova克拉约瓦 feud.Barony = &克拉约瓦CraiovaBarony{}

func init() {
    f := BCraiova克拉约瓦.(*克拉约瓦CraiovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "craiova",
		TitleName: "克拉约瓦",
		TitleCode: "b_craiova",
	}
}
