package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉塔KarataBarony struct {
	feud.BaseBarony
}

var BKarata卡拉塔 feud.Barony = &卡拉塔KarataBarony{}

func init() {
    f := BKarata卡拉塔.(*卡拉塔KarataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karata",
		TitleName: "卡拉塔",
		TitleCode: "b_karata",
	}
}
