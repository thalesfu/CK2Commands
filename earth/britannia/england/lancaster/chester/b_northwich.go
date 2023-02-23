package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺斯维奇NorthwichBarony struct {
	feud.BaseBarony
}

var BNorthwich诺斯维奇 feud.Barony = &诺斯维奇NorthwichBarony{}

func init() {
	f := BNorthwich诺斯维奇.(*诺斯维奇NorthwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "northwich",
		TitleName: "诺斯维奇",
		TitleCode: "b_northwich",
	}
}
