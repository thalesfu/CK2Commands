package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科波利斯AnatolnikopolisBarony struct {
	feud.BaseBarony
}

var BAnatolnikopolis尼科波利斯 feud.Barony = &尼科波利斯AnatolnikopolisBarony{}

func init() {
	f := BAnatolnikopolis尼科波利斯.(*尼科波利斯AnatolnikopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anatolnikopolis",
		TitleName: "尼科波利斯",
		TitleCode: "b_anatolnikopolis",
	}
}
