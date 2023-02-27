package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别洛亚尔BeloyarskiyBarony struct {
	feud.BaseBarony
}

var BBeloyarskiy别洛亚尔 feud.Barony = &别洛亚尔BeloyarskiyBarony{}

func init() {
    f := BBeloyarskiy别洛亚尔.(*别洛亚尔BeloyarskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beloyarskiy",
		TitleName: "别洛亚尔",
		TitleCode: "b_beloyarskiy",
	}
}
