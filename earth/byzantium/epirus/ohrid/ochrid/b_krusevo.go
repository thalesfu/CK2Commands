package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁舍沃KrusevoBarony struct {
	feud.BaseBarony
}

var BKrusevo克鲁舍沃 feud.Barony = &克鲁舍沃KrusevoBarony{}

func init() {
	f := BKrusevo克鲁舍沃.(*克鲁舍沃KrusevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krusevo",
		TitleName: "克鲁舍沃",
		TitleCode: "b_krusevo",
	}
}
