package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡布德拉汉格KabudrahangBarony struct {
	feud.BaseBarony
}

var BKabudrahang卡布德拉汉格 feud.Barony = &卡布德拉汉格KabudrahangBarony{}

func init() {
	f := BKabudrahang卡布德拉汉格.(*卡布德拉汉格KabudrahangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabudrahang",
		TitleName: "卡布德拉汉格",
		TitleCode: "b_kabudrahang",
	}
}
