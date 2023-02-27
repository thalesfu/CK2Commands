package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣让St_jeanBarony struct {
	feud.BaseBarony
}

var BSt_jean圣让 feud.Barony = &圣让St_jeanBarony{}

func init() {
    f := BSt_jean圣让.(*圣让St_jeanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_jean",
		TitleName: "圣让",
		TitleCode: "b_st_jean",
	}
}
