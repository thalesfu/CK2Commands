package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶弗林YeaveringBarony struct {
	feud.BaseBarony
}

var BYeavering耶弗林 feud.Barony = &耶弗林YeaveringBarony{}

func init() {
	f := BYeavering耶弗林.(*耶弗林YeaveringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeavering",
		TitleName: "耶弗林",
		TitleCode: "b_yeavering",
	}
}
