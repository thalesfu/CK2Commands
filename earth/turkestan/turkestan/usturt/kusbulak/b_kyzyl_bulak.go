package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒布拉克Kyzyl_bulakBarony struct {
	feud.BaseBarony
}

var BKyzyl_bulak克孜勒布拉克 feud.Barony = &克孜勒布拉克Kyzyl_bulakBarony{}

func init() {
    f := BKyzyl_bulak克孜勒布拉克.(*克孜勒布拉克Kyzyl_bulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_bulak",
		TitleName: "克孜勒布拉克",
		TitleCode: "b_kyzyl_bulak",
	}
}
