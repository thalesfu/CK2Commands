package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉米塔KalamitaBarony struct {
	feud.BaseBarony
}

var BKalamita卡拉米塔 feud.Barony = &卡拉米塔KalamitaBarony{}

func init() {
    f := BKalamita卡拉米塔.(*卡拉米塔KalamitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalamita",
		TitleName: "卡拉米塔",
		TitleCode: "b_kalamita",
	}
}
