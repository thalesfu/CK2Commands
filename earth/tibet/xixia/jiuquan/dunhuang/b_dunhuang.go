package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦煌DunhuangBarony struct {
	feud.BaseBarony
}

var BDunhuang敦煌 feud.Barony = &敦煌DunhuangBarony{}

func init() {
    f := BDunhuang敦煌.(*敦煌DunhuangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunhuang",
		TitleName: "敦煌",
		TitleCode: "b_dunhuang",
	}
}
