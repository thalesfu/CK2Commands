package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希梅纳德拉弗龙特拉JimenadelafronteraBarony struct {
	feud.BaseBarony
}

var BJimenadelafrontera希梅纳德拉弗龙特拉 feud.Barony = &希梅纳德拉弗龙特拉JimenadelafronteraBarony{}

func init() {
	f := BJimenadelafrontera希梅纳德拉弗龙特拉.(*希梅纳德拉弗龙特拉JimenadelafronteraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jimenadelafrontera",
		TitleName: "希梅纳德拉弗龙特拉",
		TitleCode: "b_jimenadelafrontera",
	}
}
