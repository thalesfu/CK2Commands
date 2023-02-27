package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卡尔KalkalBarony struct {
	feud.BaseBarony
}

var BKalkal卡尔卡尔 feud.Barony = &卡尔卡尔KalkalBarony{}

func init() {
    f := BKalkal卡尔卡尔.(*卡尔卡尔KalkalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalkal",
		TitleName: "卡尔卡尔",
		TitleCode: "b_kalkal",
	}
}
