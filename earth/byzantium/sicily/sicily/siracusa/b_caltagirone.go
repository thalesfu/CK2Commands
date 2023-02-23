package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔吉罗内CaltagironeBarony struct {
	feud.BaseBarony
}

var BCaltagirone卡尔塔吉罗内 feud.Barony = &卡尔塔吉罗内CaltagironeBarony{}

func init() {
	f := BCaltagirone卡尔塔吉罗内.(*卡尔塔吉罗内CaltagironeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caltagirone",
		TitleName: "卡尔塔吉罗内",
		TitleCode: "b_caltagirone",
	}
}
