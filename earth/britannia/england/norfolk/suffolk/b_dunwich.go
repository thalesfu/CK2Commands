package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓尼奇DunwichBarony struct {
	feud.BaseBarony
}

var BDunwich邓尼奇 feud.Barony = &邓尼奇DunwichBarony{}

func init() {
	f := BDunwich邓尼奇.(*邓尼奇DunwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunwich",
		TitleName: "邓尼奇",
		TitleCode: "b_dunwich",
	}
}
