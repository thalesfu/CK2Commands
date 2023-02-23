package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃那俱尼HanakuniBarony struct {
	feud.BaseBarony
}

var BHanakuni诃那俱尼 feud.Barony = &诃那俱尼HanakuniBarony{}

func init() {
	f := BHanakuni诃那俱尼.(*诃那俱尼HanakuniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hanakuni",
		TitleName: "诃那俱尼",
		TitleCode: "b_hanakuni",
	}
}
