package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔塔瓦拉SortavalaBarony struct {
	feud.BaseBarony
}

var BSortavala索尔塔瓦拉 feud.Barony = &索尔塔瓦拉SortavalaBarony{}

func init() {
    f := BSortavala索尔塔瓦拉.(*索尔塔瓦拉SortavalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sortavala",
		TitleName: "索尔塔瓦拉",
		TitleCode: "b_sortavala",
	}
}
