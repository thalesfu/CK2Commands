package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马蒂尼MartignyBarony struct {
	feud.BaseBarony
}

var BMartigny马蒂尼 feud.Barony = &马蒂尼MartignyBarony{}

func init() {
	f := BMartigny马蒂尼.(*马蒂尼MartignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martigny",
		TitleName: "马蒂尼",
		TitleCode: "b_martigny",
	}
}
