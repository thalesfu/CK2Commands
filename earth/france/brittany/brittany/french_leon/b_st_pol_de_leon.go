package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣波勒德莱昂St_pol_de_leonBarony struct {
	feud.BaseBarony
}

var BSt_pol_de_leon圣波勒德莱昂 feud.Barony = &圣波勒德莱昂St_pol_de_leonBarony{}

func init() {
    f := BSt_pol_de_leon圣波勒德莱昂.(*圣波勒德莱昂St_pol_de_leonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_pol_de_leon",
		TitleName: "圣波勒德莱昂",
		TitleCode: "b_st_pol_de_leon",
	}
}
