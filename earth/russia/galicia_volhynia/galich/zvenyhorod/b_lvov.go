package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利沃夫LvovBarony struct {
	feud.BaseBarony
}

var BLvov利沃夫 feud.Barony = &利沃夫LvovBarony{}

func init() {
    f := BLvov利沃夫.(*利沃夫LvovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lvov",
		TitleName: "利沃夫",
		TitleCode: "b_lvov",
	}
}
