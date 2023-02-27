package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里加维BalligaviBarony struct {
	feud.BaseBarony
}

var BBalligavi巴里加维 feud.Barony = &巴里加维BalligaviBarony{}

func init() {
    f := BBalligavi巴里加维.(*巴里加维BalligaviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balligavi",
		TitleName: "巴里加维",
		TitleCode: "b_balligavi",
	}
}
