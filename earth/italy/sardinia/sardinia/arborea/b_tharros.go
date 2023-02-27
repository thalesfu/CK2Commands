package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔罗斯TharrosBarony struct {
	feud.BaseBarony
}

var BTharros塔罗斯 feud.Barony = &塔罗斯TharrosBarony{}

func init() {
    f := BTharros塔罗斯.(*塔罗斯TharrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tharros",
		TitleName: "塔罗斯",
		TitleCode: "b_tharros",
	}
}
