package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维杰瓦诺VigevanoBarony struct {
	feud.BaseBarony
}

var BVigevano维杰瓦诺 feud.Barony = &维杰瓦诺VigevanoBarony{}

func init() {
	f := BVigevano维杰瓦诺.(*维杰瓦诺VigevanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vigevano",
		TitleName: "维杰瓦诺",
		TitleCode: "b_vigevano",
	}
}
