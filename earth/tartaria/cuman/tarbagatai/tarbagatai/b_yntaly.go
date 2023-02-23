package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩塔雷YntalyBarony struct {
	feud.BaseBarony
}

var BYntaly恩塔雷 feud.Barony = &恩塔雷YntalyBarony{}

func init() {
	f := BYntaly恩塔雷.(*恩塔雷YntalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yntaly",
		TitleName: "恩塔雷",
		TitleCode: "b_yntaly",
	}
}
