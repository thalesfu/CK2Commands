package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔BourgBarony struct {
	feud.BaseBarony
}

var BBourg布尔 feud.Barony = &布尔BourgBarony{}

func init() {
	f := BBourg布尔.(*布尔BourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourg",
		TitleName: "布尔",
		TitleCode: "b_bourg",
	}
}
