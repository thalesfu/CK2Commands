package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔巴斯亚DarbasiyahBarony struct {
	feud.BaseBarony
}

var BDarbasiyah达尔巴斯亚 feud.Barony = &达尔巴斯亚DarbasiyahBarony{}

func init() {
    f := BDarbasiyah达尔巴斯亚.(*达尔巴斯亚DarbasiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darbasiyah",
		TitleName: "达尔巴斯亚",
		TitleCode: "b_darbasiyah",
	}
}
