package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔曼KilmunBarony struct {
	feud.BaseBarony
}

var BKilmun基尔曼 feud.Barony = &基尔曼KilmunBarony{}

func init() {
    f := BKilmun基尔曼.(*基尔曼KilmunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilmun",
		TitleName: "基尔曼",
		TitleCode: "b_kilmun",
	}
}
