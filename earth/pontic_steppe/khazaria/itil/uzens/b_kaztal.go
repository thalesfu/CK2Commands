package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡兹塔勒KaztalBarony struct {
	feud.BaseBarony
}

var BKaztal卡兹塔勒 feud.Barony = &卡兹塔勒KaztalBarony{}

func init() {
    f := BKaztal卡兹塔勒.(*卡兹塔勒KaztalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaztal",
		TitleName: "卡兹塔勒",
		TitleCode: "b_kaztal",
	}
}
