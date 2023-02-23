package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娘拉NianglaBarony struct {
	feud.BaseBarony
}

var BNiangla娘拉 feud.Barony = &娘拉NianglaBarony{}

func init() {
	f := BNiangla娘拉.(*娘拉NianglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niangla",
		TitleName: "娘拉",
		TitleCode: "b_niangla",
	}
}
