package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉嘎瓦LagawaBarony struct {
	feud.BaseBarony
}

var BLagawa拉嘎瓦 feud.Barony = &拉嘎瓦LagawaBarony{}

func init() {
    f := BLagawa拉嘎瓦.(*拉嘎瓦LagawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lagawa",
		TitleName: "拉嘎瓦",
		TitleCode: "b_lagawa",
	}
}
