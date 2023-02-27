package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯_布拉克KosbulakBarony struct {
	feud.BaseBarony
}

var BKosbulak科斯_布拉克 feud.Barony = &科斯_布拉克KosbulakBarony{}

func init() {
    f := BKosbulak科斯_布拉克.(*科斯_布拉克KosbulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosbulak",
		TitleName: "科斯_布拉克",
		TitleCode: "b_kosbulak",
	}
}
