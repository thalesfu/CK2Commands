package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔格塔巴扎尔TagtabazarBarony struct {
	feud.BaseBarony
}

var BTagtabazar塔格塔巴扎尔 feud.Barony = &塔格塔巴扎尔TagtabazarBarony{}

func init() {
	f := BTagtabazar塔格塔巴扎尔.(*塔格塔巴扎尔TagtabazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagtabazar",
		TitleName: "塔格塔巴扎尔",
		TitleCode: "b_tagtabazar",
	}
}
