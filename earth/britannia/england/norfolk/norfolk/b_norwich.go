package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺里奇NorwichBarony struct {
	feud.BaseBarony
}

var BNorwich诺里奇 feud.Barony = &诺里奇NorwichBarony{}

func init() {
	f := BNorwich诺里奇.(*诺里奇NorwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norwich",
		TitleName: "诺里奇",
		TitleCode: "b_norwich",
	}
}
