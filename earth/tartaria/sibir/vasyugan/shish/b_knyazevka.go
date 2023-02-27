package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼亚泽夫卡KnyazevkaBarony struct {
	feud.BaseBarony
}

var BKnyazevka克尼亚泽夫卡 feud.Barony = &克尼亚泽夫卡KnyazevkaBarony{}

func init() {
    f := BKnyazevka克尼亚泽夫卡.(*克尼亚泽夫卡KnyazevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knyazevka",
		TitleName: "克尼亚泽夫卡",
		TitleCode: "b_knyazevka",
	}
}
