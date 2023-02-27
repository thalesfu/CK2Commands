package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰波日尼亚LampozhnyaBarony struct {
	feud.BaseBarony
}

var BLampozhnya兰波日尼亚 feud.Barony = &兰波日尼亚LampozhnyaBarony{}

func init() {
    f := BLampozhnya兰波日尼亚.(*兰波日尼亚LampozhnyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lampozhnya",
		TitleName: "兰波日尼亚",
		TitleCode: "b_lampozhnya",
	}
}
