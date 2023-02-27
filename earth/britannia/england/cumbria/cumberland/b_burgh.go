package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒BurghBarony struct {
	feud.BaseBarony
}

var BBurgh巴勒 feud.Barony = &巴勒BurghBarony{}

func init() {
    f := BBurgh巴勒.(*巴勒BurghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgh",
		TitleName: "巴勒",
		TitleCode: "b_burgh",
	}
}
