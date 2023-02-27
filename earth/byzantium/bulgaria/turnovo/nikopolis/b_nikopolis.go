package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科波利斯NikopolisBarony struct {
	feud.BaseBarony
}

var BNikopolis尼科波利斯 feud.Barony = &尼科波利斯NikopolisBarony{}

func init() {
    f := BNikopolis尼科波利斯.(*尼科波利斯NikopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikopolis",
		TitleName: "尼科波利斯",
		TitleCode: "b_nikopolis",
	}
}
