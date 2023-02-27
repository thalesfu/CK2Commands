package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科波尔NikopolBarony struct {
	feud.BaseBarony
}

var BNikopol尼科波尔 feud.Barony = &尼科波尔NikopolBarony{}

func init() {
    f := BNikopol尼科波尔.(*尼科波尔NikopolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikopol",
		TitleName: "尼科波尔",
		TitleCode: "b_nikopol",
	}
}
