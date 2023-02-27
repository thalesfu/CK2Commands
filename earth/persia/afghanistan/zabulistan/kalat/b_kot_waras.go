package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库特瓦拉斯Kot_warasBarony struct {
	feud.BaseBarony
}

var BKot_waras库特瓦拉斯 feud.Barony = &库特瓦拉斯Kot_warasBarony{}

func init() {
    f := BKot_waras库特瓦拉斯.(*库特瓦拉斯Kot_warasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kot_waras",
		TitleName: "库特瓦拉斯",
		TitleCode: "b_kot_waras",
	}
}
