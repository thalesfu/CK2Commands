package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉万萨尔RavansarBarony struct {
	feud.BaseBarony
}

var BRavansar拉万萨尔 feud.Barony = &拉万萨尔RavansarBarony{}

func init() {
	f := BRavansar拉万萨尔.(*拉万萨尔RavansarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravansar",
		TitleName: "拉万萨尔",
		TitleCode: "b_ravansar",
	}
}
