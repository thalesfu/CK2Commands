package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别拉亚BelayaBarony struct {
	feud.BaseBarony
}

var BBelaya别拉亚 feud.Barony = &别拉亚BelayaBarony{}

func init() {
    f := BBelaya别拉亚.(*别拉亚BelayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belaya",
		TitleName: "别拉亚",
		TitleCode: "b_belaya",
	}
}
