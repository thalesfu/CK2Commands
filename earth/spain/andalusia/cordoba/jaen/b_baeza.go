package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴埃萨BaezaBarony struct {
	feud.BaseBarony
}

var BBaeza巴埃萨 feud.Barony = &巴埃萨BaezaBarony{}

func init() {
    f := BBaeza巴埃萨.(*巴埃萨BaezaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baeza",
		TitleName: "巴埃萨",
		TitleCode: "b_baeza",
	}
}
