package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗利亚KhorolyaBarony struct {
	feud.BaseBarony
}

var BKhorolya霍罗利亚 feud.Barony = &霍罗利亚KhorolyaBarony{}

func init() {
    f := BKhorolya霍罗利亚.(*霍罗利亚KhorolyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorolya",
		TitleName: "霍罗利亚",
		TitleCode: "b_khorolya",
	}
}
