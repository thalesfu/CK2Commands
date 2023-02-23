package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔堡BourbourgBarony struct {
	feud.BaseBarony
}

var BBourbourg布尔堡 feud.Barony = &布尔堡BourbourgBarony{}

func init() {
	f := BBourbourg布尔堡.(*布尔堡BourbourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bourbourg",
		TitleName: "布尔堡",
		TitleCode: "b_bourbourg",
	}
}
