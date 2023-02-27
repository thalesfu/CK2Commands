package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡盘KapanBarony struct {
	feud.BaseBarony
}

var BKapan卡盘 feud.Barony = &卡盘KapanBarony{}

func init() {
    f := BKapan卡盘.(*卡盘KapanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapan",
		TitleName: "卡盘",
		TitleCode: "b_kapan",
	}
}
