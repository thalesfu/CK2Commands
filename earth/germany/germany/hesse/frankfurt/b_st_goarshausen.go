package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣戈阿斯豪森St_goarshausenBarony struct {
	feud.BaseBarony
}

var BSt_goarshausen圣戈阿斯豪森 feud.Barony = &圣戈阿斯豪森St_goarshausenBarony{}

func init() {
    f := BSt_goarshausen圣戈阿斯豪森.(*圣戈阿斯豪森St_goarshausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_goarshausen",
		TitleName: "圣戈阿斯豪森",
		TitleCode: "b_st_goarshausen",
	}
}
