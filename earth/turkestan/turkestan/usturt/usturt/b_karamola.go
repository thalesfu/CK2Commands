package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉莫拉KaramolaBarony struct {
	feud.BaseBarony
}

var BKaramola卡拉莫拉 feud.Barony = &卡拉莫拉KaramolaBarony{}

func init() {
	f := BKaramola卡拉莫拉.(*卡拉莫拉KaramolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karamola",
		TitleName: "卡拉莫拉",
		TitleCode: "b_karamola",
	}
}
