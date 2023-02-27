package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡哈纳拉卡拉KhannalkalaBarony struct {
	feud.BaseBarony
}

var BKhannalkala卡哈纳拉卡拉 feud.Barony = &卡哈纳拉卡拉KhannalkalaBarony{}

func init() {
    f := BKhannalkala卡哈纳拉卡拉.(*卡哈纳拉卡拉KhannalkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khannalkala",
		TitleName: "卡哈纳拉卡拉",
		TitleCode: "b_khannalkala",
	}
}
