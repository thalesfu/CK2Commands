package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷克_乌鲁KrykuruBarony struct {
	feud.BaseBarony
}

var BKrykuru克雷克_乌鲁 feud.Barony = &克雷克_乌鲁KrykuruBarony{}

func init() {
    f := BKrykuru克雷克_乌鲁.(*克雷克_乌鲁KrykuruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krykuru",
		TitleName: "克雷克_乌鲁",
		TitleCode: "b_krykuru",
	}
}
