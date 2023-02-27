package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗萨DzongsarBarony struct {
	feud.BaseBarony
}

var BDzongsar宗萨 feud.Barony = &宗萨DzongsarBarony{}

func init() {
    f := BDzongsar宗萨.(*宗萨DzongsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzongsar",
		TitleName: "宗萨",
		TitleCode: "b_dzongsar",
	}
}
