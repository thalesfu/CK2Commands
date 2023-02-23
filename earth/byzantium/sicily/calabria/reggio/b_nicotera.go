package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科泰拉NicoteraBarony struct {
	feud.BaseBarony
}

var BNicotera尼科泰拉 feud.Barony = &尼科泰拉NicoteraBarony{}

func init() {
	f := BNicotera尼科泰拉.(*尼科泰拉NicoteraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nicotera",
		TitleName: "尼科泰拉",
		TitleCode: "b_nicotera",
	}
}
