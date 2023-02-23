package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯塔利夫IstalifBarony struct {
	feud.BaseBarony
}

var BIstalif伊斯塔利夫 feud.Barony = &伊斯塔利夫IstalifBarony{}

func init() {
	f := BIstalif伊斯塔利夫.(*伊斯塔利夫IstalifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "istalif",
		TitleName: "伊斯塔利夫",
		TitleCode: "b_istalif",
	}
}
