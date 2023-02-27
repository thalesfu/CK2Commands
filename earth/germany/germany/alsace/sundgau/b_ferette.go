package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费雷特FeretteBarony struct {
	feud.BaseBarony
}

var BFerette费雷特 feud.Barony = &费雷特FeretteBarony{}

func init() {
    f := BFerette费雷特.(*费雷特FeretteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferette",
		TitleName: "费雷特",
		TitleCode: "b_ferette",
	}
}
