package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞利祖尔SellizureBarony struct {
	feud.BaseBarony
}

var BSellizure塞利祖尔 feud.Barony = &塞利祖尔SellizureBarony{}

func init() {
    f := BSellizure塞利祖尔.(*塞利祖尔SellizureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sellizure",
		TitleName: "塞利祖尔",
		TitleCode: "b_sellizure",
	}
}
