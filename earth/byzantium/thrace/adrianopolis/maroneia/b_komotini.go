package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科莫蒂尼KomotiniBarony struct {
	feud.BaseBarony
}

var BKomotini科莫蒂尼 feud.Barony = &科莫蒂尼KomotiniBarony{}

func init() {
    f := BKomotini科莫蒂尼.(*科莫蒂尼KomotiniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komotini",
		TitleName: "科莫蒂尼",
		TitleCode: "b_komotini",
	}
}
