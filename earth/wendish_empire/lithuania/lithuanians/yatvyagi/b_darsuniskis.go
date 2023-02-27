package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔苏尼斯基斯DarsuniskisBarony struct {
	feud.BaseBarony
}

var BDarsuniskis达尔苏尼斯基斯 feud.Barony = &达尔苏尼斯基斯DarsuniskisBarony{}

func init() {
    f := BDarsuniskis达尔苏尼斯基斯.(*达尔苏尼斯基斯DarsuniskisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darsuniskis",
		TitleName: "达尔苏尼斯基斯",
		TitleCode: "b_darsuniskis",
	}
}
