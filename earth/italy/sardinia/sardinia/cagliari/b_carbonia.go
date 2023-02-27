package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔博尼亚CarboniaBarony struct {
	feud.BaseBarony
}

var BCarbonia卡尔博尼亚 feud.Barony = &卡尔博尼亚CarboniaBarony{}

func init() {
    f := BCarbonia卡尔博尼亚.(*卡尔博尼亚CarboniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carbonia",
		TitleName: "卡尔博尼亚",
		TitleCode: "b_carbonia",
	}
}
