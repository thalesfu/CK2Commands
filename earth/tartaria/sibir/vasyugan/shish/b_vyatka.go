package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚特卡VyatkaBarony struct {
	feud.BaseBarony
}

var BVyatka维亚特卡 feud.Barony = &维亚特卡VyatkaBarony{}

func init() {
    f := BVyatka维亚特卡.(*维亚特卡VyatkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyatka",
		TitleName: "维亚特卡",
		TitleCode: "b_vyatka",
	}
}
