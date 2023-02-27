package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希利克塔ShiliktaBarony struct {
	feud.BaseBarony
}

var BShilikta希利克塔 feud.Barony = &希利克塔ShiliktaBarony{}

func init() {
    f := BShilikta希利克塔.(*希利克塔ShiliktaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shilikta",
		TitleName: "希利克塔",
		TitleCode: "b_shilikta",
	}
}
