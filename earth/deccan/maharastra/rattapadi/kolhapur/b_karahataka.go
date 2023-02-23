package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉荷塔克KarahatakaBarony struct {
	feud.BaseBarony
}

var BKarahataka卡拉荷塔克 feud.Barony = &卡拉荷塔克KarahatakaBarony{}

func init() {
	f := BKarahataka卡拉荷塔克.(*卡拉荷塔克KarahatakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karahataka",
		TitleName: "卡拉荷塔克",
		TitleCode: "b_karahataka",
	}
}
