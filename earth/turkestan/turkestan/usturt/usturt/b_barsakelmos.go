package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔萨_克尔梅斯BarsakelmosBarony struct {
	feud.BaseBarony
}

var BBarsakelmos巴尔萨_克尔梅斯 feud.Barony = &巴尔萨_克尔梅斯BarsakelmosBarony{}

func init() {
    f := BBarsakelmos巴尔萨_克尔梅斯.(*巴尔萨_克尔梅斯BarsakelmosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barsakelmos",
		TitleName: "巴尔萨_克尔梅斯",
		TitleCode: "b_barsakelmos",
	}
}
