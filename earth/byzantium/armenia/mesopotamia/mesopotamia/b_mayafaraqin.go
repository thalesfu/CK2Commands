package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛亚法拉琴MayafaraqinBarony struct {
	feud.BaseBarony
}

var BMayafaraqin玛亚法拉琴 feud.Barony = &玛亚法拉琴MayafaraqinBarony{}

func init() {
	f := BMayafaraqin玛亚法拉琴.(*玛亚法拉琴MayafaraqinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayafaraqin",
		TitleName: "玛亚法拉琴",
		TitleCode: "b_mayafaraqin",
	}
}
