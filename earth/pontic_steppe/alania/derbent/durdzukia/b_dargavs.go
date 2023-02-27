package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔加夫斯DargavsBarony struct {
	feud.BaseBarony
}

var BDargavs达尔加夫斯 feud.Barony = &达尔加夫斯DargavsBarony{}

func init() {
    f := BDargavs达尔加夫斯.(*达尔加夫斯DargavsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dargavs",
		TitleName: "达尔加夫斯",
		TitleCode: "b_dargavs",
	}
}
