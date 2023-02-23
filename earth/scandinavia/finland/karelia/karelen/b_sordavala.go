package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔塔瓦拉SordavalaBarony struct {
	feud.BaseBarony
}

var BSordavala索尔塔瓦拉 feud.Barony = &索尔塔瓦拉SordavalaBarony{}

func init() {
	f := BSordavala索尔塔瓦拉.(*索尔塔瓦拉SordavalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sordavala",
		TitleName: "索尔塔瓦拉",
		TitleCode: "b_sordavala",
	}
}
