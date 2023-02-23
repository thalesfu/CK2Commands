package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫巴RahbahBarony struct {
	feud.BaseBarony
}

var BRahbah拉赫巴 feud.Barony = &拉赫巴RahbahBarony{}

func init() {
	f := BRahbah拉赫巴.(*拉赫巴RahbahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rahbah",
		TitleName: "拉赫巴",
		TitleCode: "b_rahbah",
	}
}
