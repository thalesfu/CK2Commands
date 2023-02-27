package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴胡BahuBarony struct {
	feud.BaseBarony
}

var BBahu巴胡 feud.Barony = &巴胡BahuBarony{}

func init() {
    f := BBahu巴胡.(*巴胡BahuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahu",
		TitleName: "巴胡",
		TitleCode: "b_bahu",
	}
}
