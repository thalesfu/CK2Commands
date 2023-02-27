package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉索尔KarasorBarony struct {
	feud.BaseBarony
}

var BKarasor卡拉索尔 feud.Barony = &卡拉索尔KarasorBarony{}

func init() {
    f := BKarasor卡拉索尔.(*卡拉索尔KarasorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karasor",
		TitleName: "卡拉索尔",
		TitleCode: "b_karasor",
	}
}
