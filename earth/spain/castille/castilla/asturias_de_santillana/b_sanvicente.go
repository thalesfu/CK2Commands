package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣维森特SanvicenteBarony struct {
	feud.BaseBarony
}

var BSanvicente圣维森特 feud.Barony = &圣维森特SanvicenteBarony{}

func init() {
    f := BSanvicente圣维森特.(*圣维森特SanvicenteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanvicente",
		TitleName: "圣维森特",
		TitleCode: "b_sanvicente",
	}
}
