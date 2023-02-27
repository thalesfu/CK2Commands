package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉塔尼亚索尔CalatanazorBarony struct {
	feud.BaseBarony
}

var BCalatanazor卡拉塔尼亚索尔 feud.Barony = &卡拉塔尼亚索尔CalatanazorBarony{}

func init() {
    f := BCalatanazor卡拉塔尼亚索尔.(*卡拉塔尼亚索尔CalatanazorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calatanazor",
		TitleName: "卡拉塔尼亚索尔",
		TitleCode: "b_calatanazor",
	}
}
