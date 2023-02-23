package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴朗戈卡BalangokaBarony struct {
	feud.BaseBarony
}

var BBalangoka巴朗戈卡 feud.Barony = &巴朗戈卡BalangokaBarony{}

func init() {
	f := BBalangoka巴朗戈卡.(*巴朗戈卡BalangokaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balangoka",
		TitleName: "巴朗戈卡",
		TitleCode: "b_balangoka",
	}
}
