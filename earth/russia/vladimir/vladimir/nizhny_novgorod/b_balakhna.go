package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉赫纳BalakhnaBarony struct {
	feud.BaseBarony
}

var BBalakhna巴拉赫纳 feud.Barony = &巴拉赫纳BalakhnaBarony{}

func init() {
    f := BBalakhna巴拉赫纳.(*巴拉赫纳BalakhnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balakhna",
		TitleName: "巴拉赫纳",
		TitleCode: "b_balakhna",
	}
}
