package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯佩CaspeBarony struct {
	feud.BaseBarony
}

var BCaspe卡斯佩 feud.Barony = &卡斯佩CaspeBarony{}

func init() {
	f := BCaspe卡斯佩.(*卡斯佩CaspeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caspe",
		TitleName: "卡斯佩",
		TitleCode: "b_caspe",
	}
}
