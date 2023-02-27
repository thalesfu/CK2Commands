package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔特拉BartellaBarony struct {
	feud.BaseBarony
}

var BBartella巴尔特拉 feud.Barony = &巴尔特拉BartellaBarony{}

func init() {
    f := BBartella巴尔特拉.(*巴尔特拉BartellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bartella",
		TitleName: "巴尔特拉",
		TitleCode: "b_bartella",
	}
}
