package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉卡尔KarakalBarony struct {
	feud.BaseBarony
}

var BKarakal卡拉卡尔 feud.Barony = &卡拉卡尔KarakalBarony{}

func init() {
	f := BKarakal卡拉卡尔.(*卡拉卡尔KarakalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karakal",
		TitleName: "卡拉卡尔",
		TitleCode: "b_karakal",
	}
}
