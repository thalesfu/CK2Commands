package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣但尼StdenisBarony struct {
	feud.BaseBarony
}

var BStdenis圣但尼 feud.Barony = &圣但尼StdenisBarony{}

func init() {
    f := BStdenis圣但尼.(*圣但尼StdenisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stdenis",
		TitleName: "圣但尼",
		TitleCode: "b_stdenis",
	}
}
