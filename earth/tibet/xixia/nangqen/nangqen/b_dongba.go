package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东坝DongbaBarony struct {
	feud.BaseBarony
}

var BDongba东坝 feud.Barony = &东坝DongbaBarony{}

func init() {
	f := BDongba东坝.(*东坝DongbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dongba",
		TitleName: "东坝",
		TitleCode: "b_dongba",
	}
}
