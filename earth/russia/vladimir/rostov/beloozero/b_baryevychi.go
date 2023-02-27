package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里耶维奇BaryevychiBarony struct {
	feud.BaseBarony
}

var BBaryevychi巴里耶维奇 feud.Barony = &巴里耶维奇BaryevychiBarony{}

func init() {
    f := BBaryevychi巴里耶维奇.(*巴里耶维奇BaryevychiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baryevychi",
		TitleName: "巴里耶维奇",
		TitleCode: "b_baryevychi",
	}
}
