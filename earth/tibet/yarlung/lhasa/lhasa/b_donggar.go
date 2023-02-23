package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东嘎DonggarBarony struct {
	feud.BaseBarony
}

var BDonggar东嘎 feud.Barony = &东嘎DonggarBarony{}

func init() {
	f := BDonggar东嘎.(*东嘎DonggarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donggar",
		TitleName: "东嘎",
		TitleCode: "b_donggar",
	}
}
