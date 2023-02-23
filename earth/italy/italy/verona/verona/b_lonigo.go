package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛尼戈LonigoBarony struct {
	feud.BaseBarony
}

var BLonigo洛尼戈 feud.Barony = &洛尼戈LonigoBarony{}

func init() {
	f := BLonigo洛尼戈.(*洛尼戈LonigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lonigo",
		TitleName: "洛尼戈",
		TitleCode: "b_lonigo",
	}
}
