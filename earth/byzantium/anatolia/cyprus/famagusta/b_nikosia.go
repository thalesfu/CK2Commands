package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科西亚NikosiaBarony struct {
	feud.BaseBarony
}

var BNikosia尼科西亚 feud.Barony = &尼科西亚NikosiaBarony{}

func init() {
    f := BNikosia尼科西亚.(*尼科西亚NikosiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikosia",
		TitleName: "尼科西亚",
		TitleCode: "b_nikosia",
	}
}
