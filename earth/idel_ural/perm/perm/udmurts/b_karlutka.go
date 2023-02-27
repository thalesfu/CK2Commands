package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔卢特卡KarlutkaBarony struct {
	feud.BaseBarony
}

var BKarlutka卡尔卢特卡 feud.Barony = &卡尔卢特卡KarlutkaBarony{}

func init() {
    f := BKarlutka卡尔卢特卡.(*卡尔卢特卡KarlutkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karlutka",
		TitleName: "卡尔卢特卡",
		TitleCode: "b_karlutka",
	}
}
