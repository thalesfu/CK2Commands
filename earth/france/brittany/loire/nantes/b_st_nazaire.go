package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣纳泽尔St_nazaireBarony struct {
	feud.BaseBarony
}

var BSt_nazaire圣纳泽尔 feud.Barony = &圣纳泽尔St_nazaireBarony{}

func init() {
    f := BSt_nazaire圣纳泽尔.(*圣纳泽尔St_nazaireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_nazaire",
		TitleName: "圣纳泽尔",
		TitleCode: "b_st_nazaire",
	}
}
