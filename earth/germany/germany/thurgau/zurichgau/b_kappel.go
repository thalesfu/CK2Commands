package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡珀尔KappelBarony struct {
	feud.BaseBarony
}

var BKappel卡珀尔 feud.Barony = &卡珀尔KappelBarony{}

func init() {
    f := BKappel卡珀尔.(*卡珀尔KappelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kappel",
		TitleName: "卡珀尔",
		TitleCode: "b_kappel",
	}
}
