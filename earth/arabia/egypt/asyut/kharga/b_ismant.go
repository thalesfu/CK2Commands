package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾斯曼IsmantBarony struct {
	feud.BaseBarony
}

var BIsmant艾斯曼 feud.Barony = &艾斯曼IsmantBarony{}

func init() {
    f := BIsmant艾斯曼.(*艾斯曼IsmantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ismant",
		TitleName: "艾斯曼",
		TitleCode: "b_ismant",
	}
}
