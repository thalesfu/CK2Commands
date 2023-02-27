package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰希费El_jehifaBarony struct {
	feud.BaseBarony
}

var BEl_jehifa杰希费 feud.Barony = &杰希费El_jehifaBarony{}

func init() {
    f := BEl_jehifa杰希费.(*杰希费El_jehifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_jehifa",
		TitleName: "杰希费",
		TitleCode: "b_el_jehifa",
	}
}
