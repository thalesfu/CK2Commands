package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤登堡JudenburgBarony struct {
	feud.BaseBarony
}

var BJudenburg尤登堡 feud.Barony = &尤登堡JudenburgBarony{}

func init() {
    f := BJudenburg尤登堡.(*尤登堡JudenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "judenburg",
		TitleName: "尤登堡",
		TitleCode: "b_judenburg",
	}
}
