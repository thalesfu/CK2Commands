package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南湖Nanhu_tarimBarony struct {
	feud.BaseBarony
}

var BNanhu_tarim南湖 feud.Barony = &南湖Nanhu_tarimBarony{}

func init() {
    f := BNanhu_tarim南湖.(*南湖Nanhu_tarimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanhu_tarim",
		TitleName: "南湖",
		TitleCode: "b_nanhu_tarim",
	}
}
