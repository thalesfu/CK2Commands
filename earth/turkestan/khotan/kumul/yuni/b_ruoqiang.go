package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婼羌RuoqiangBarony struct {
	feud.BaseBarony
}

var BRuoqiang婼羌 feud.Barony = &婼羌RuoqiangBarony{}

func init() {
    f := BRuoqiang婼羌.(*婼羌RuoqiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruoqiang",
		TitleName: "婼羌",
		TitleCode: "b_ruoqiang",
	}
}
