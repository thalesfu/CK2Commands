package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡里尼什ScarinishBarony struct {
	feud.BaseBarony
}

var BScarinish斯卡里尼什 feud.Barony = &斯卡里尼什ScarinishBarony{}

func init() {
    f := BScarinish斯卡里尼什.(*斯卡里尼什ScarinishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scarinish",
		TitleName: "斯卡里尼什",
		TitleCode: "b_scarinish",
	}
}
