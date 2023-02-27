package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维利亚卡VilakaBarony struct {
	feud.BaseBarony
}

var BVilaka维利亚卡 feud.Barony = &维利亚卡VilakaBarony{}

func init() {
    f := BVilaka维利亚卡.(*维利亚卡VilakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilaka",
		TitleName: "维利亚卡",
		TitleCode: "b_vilaka",
	}
}
