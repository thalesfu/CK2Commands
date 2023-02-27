package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡尔塔比HultabyBarony struct {
	feud.BaseBarony
}

var BHultaby胡尔塔比 feud.Barony = &胡尔塔比HultabyBarony{}

func init() {
    f := BHultaby胡尔塔比.(*胡尔塔比HultabyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hultaby",
		TitleName: "胡尔塔比",
		TitleCode: "b_hultaby",
	}
}
