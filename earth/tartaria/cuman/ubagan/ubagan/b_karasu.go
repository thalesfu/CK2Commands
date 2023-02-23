package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉苏KarasuBarony struct {
	feud.BaseBarony
}

var BKarasu卡拉苏 feud.Barony = &卡拉苏KarasuBarony{}

func init() {
	f := BKarasu卡拉苏.(*卡拉苏KarasuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karasu",
		TitleName: "卡拉苏",
		TitleCode: "b_karasu",
	}
}
