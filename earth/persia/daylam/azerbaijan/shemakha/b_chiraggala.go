package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇拉赫_卡拉ChiraggalaBarony struct {
	feud.BaseBarony
}

var BChiraggala奇拉赫_卡拉 feud.Barony = &奇拉赫_卡拉ChiraggalaBarony{}

func init() {
    f := BChiraggala奇拉赫_卡拉.(*奇拉赫_卡拉ChiraggalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiraggala",
		TitleName: "奇拉赫_卡拉",
		TitleCode: "b_chiraggala",
	}
}
