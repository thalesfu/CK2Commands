package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉绍夫BalashovBarony struct {
	feud.BaseBarony
}

var BBalashov巴拉绍夫 feud.Barony = &巴拉绍夫BalashovBarony{}

func init() {
    f := BBalashov巴拉绍夫.(*巴拉绍夫BalashovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balashov",
		TitleName: "巴拉绍夫",
		TitleCode: "b_balashov",
	}
}
