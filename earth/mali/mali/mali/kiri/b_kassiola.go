package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡西奥拉KassiolaBarony struct {
	feud.BaseBarony
}

var BKassiola卡西奥拉 feud.Barony = &卡西奥拉KassiolaBarony{}

func init() {
    f := BKassiola卡西奥拉.(*卡西奥拉KassiolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kassiola",
		TitleName: "卡西奥拉",
		TitleCode: "b_kassiola",
	}
}
