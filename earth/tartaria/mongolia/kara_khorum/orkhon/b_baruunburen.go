package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴伦布兰BaruunburenBarony struct {
	feud.BaseBarony
}

var BBaruunburen巴伦布兰 feud.Barony = &巴伦布兰BaruunburenBarony{}

func init() {
    f := BBaruunburen巴伦布兰.(*巴伦布兰BaruunburenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baruunburen",
		TitleName: "巴伦布兰",
		TitleCode: "b_baruunburen",
	}
}
