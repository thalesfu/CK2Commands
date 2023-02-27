package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡宁KaninBarony struct {
	feud.BaseBarony
}

var BKanin卡宁 feud.Barony = &卡宁KaninBarony{}

func init() {
    f := BKanin卡宁.(*卡宁KaninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanin",
		TitleName: "卡宁",
		TitleCode: "b_kanin",
	}
}
