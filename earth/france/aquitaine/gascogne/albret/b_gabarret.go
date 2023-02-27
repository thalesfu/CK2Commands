package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加巴雷GabarretBarony struct {
	feud.BaseBarony
}

var BGabarret加巴雷 feud.Barony = &加巴雷GabarretBarony{}

func init() {
    f := BGabarret加巴雷.(*加巴雷GabarretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabarret",
		TitleName: "加巴雷",
		TitleCode: "b_gabarret",
	}
}
