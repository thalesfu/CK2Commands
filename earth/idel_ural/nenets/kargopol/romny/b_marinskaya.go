package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林斯卡亚MarinskayaBarony struct {
	feud.BaseBarony
}

var BMarinskaya马林斯卡亚 feud.Barony = &马林斯卡亚MarinskayaBarony{}

func init() {
    f := BMarinskaya马林斯卡亚.(*马林斯卡亚MarinskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marinskaya",
		TitleName: "马林斯卡亚",
		TitleCode: "b_marinskaya",
	}
}
