package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔马KaarmaBarony struct {
	feud.BaseBarony
}

var BKaarma卡尔马 feud.Barony = &卡尔马KaarmaBarony{}

func init() {
    f := BKaarma卡尔马.(*卡尔马KaarmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaarma",
		TitleName: "卡尔马",
		TitleCode: "b_kaarma",
	}
}
