package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马卡特MakatBarony struct {
	feud.BaseBarony
}

var BMakat马卡特 feud.Barony = &马卡特MakatBarony{}

func init() {
	f := BMakat马卡特.(*马卡特MakatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makat",
		TitleName: "马卡特",
		TitleCode: "b_makat",
	}
}
