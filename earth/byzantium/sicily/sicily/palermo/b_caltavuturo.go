package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔武图罗CaltavuturoBarony struct {
	feud.BaseBarony
}

var BCaltavuturo卡尔塔武图罗 feud.Barony = &卡尔塔武图罗CaltavuturoBarony{}

func init() {
    f := BCaltavuturo卡尔塔武图罗.(*卡尔塔武图罗CaltavuturoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caltavuturo",
		TitleName: "卡尔塔武图罗",
		TitleCode: "b_caltavuturo",
	}
}
