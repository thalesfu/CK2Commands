package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉马塔KalamataBarony struct {
	feud.BaseBarony
}

var BKalamata卡拉马塔 feud.Barony = &卡拉马塔KalamataBarony{}

func init() {
    f := BKalamata卡拉马塔.(*卡拉马塔KalamataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalamata",
		TitleName: "卡拉马塔",
		TitleCode: "b_kalamata",
	}
}
