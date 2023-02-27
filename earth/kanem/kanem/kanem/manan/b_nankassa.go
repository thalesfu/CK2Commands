package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南卡萨NankassaBarony struct {
	feud.BaseBarony
}

var BNankassa南卡萨 feud.Barony = &南卡萨NankassaBarony{}

func init() {
    f := BNankassa南卡萨.(*南卡萨NankassaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nankassa",
		TitleName: "南卡萨",
		TitleCode: "b_nankassa",
	}
}
