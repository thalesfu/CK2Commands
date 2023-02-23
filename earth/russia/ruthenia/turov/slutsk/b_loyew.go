package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛耶夫LoyewBarony struct {
	feud.BaseBarony
}

var BLoyew洛耶夫 feud.Barony = &洛耶夫LoyewBarony{}

func init() {
	f := BLoyew洛耶夫.(*洛耶夫LoyewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loyew",
		TitleName: "洛耶夫",
		TitleCode: "b_loyew",
	}
}
