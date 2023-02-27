package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西奈塔SineitaBarony struct {
	feud.BaseBarony
}

var BSineita西奈塔 feud.Barony = &西奈塔SineitaBarony{}

func init() {
    f := BSineita西奈塔.(*西奈塔SineitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sineita",
		TitleName: "西奈塔",
		TitleCode: "b_sineita",
	}
}
