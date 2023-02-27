package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代于莱坎DaulakanBarony struct {
	feud.BaseBarony
}

var BDaulakan代于莱坎 feud.Barony = &代于莱坎DaulakanBarony{}

func init() {
    f := BDaulakan代于莱坎.(*代于莱坎DaulakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daulakan",
		TitleName: "代于莱坎",
		TitleCode: "b_daulakan",
	}
}
