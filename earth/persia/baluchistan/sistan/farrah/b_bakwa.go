package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴克瓦BakwaBarony struct {
	feud.BaseBarony
}

var BBakwa巴克瓦 feud.Barony = &巴克瓦BakwaBarony{}

func init() {
    f := BBakwa巴克瓦.(*巴克瓦BakwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakwa",
		TitleName: "巴克瓦",
		TitleCode: "b_bakwa",
	}
}
