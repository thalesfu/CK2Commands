package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛斯特堡Kloster_bergeBarony struct {
	feud.BaseBarony
}

var BKloster_berge克洛斯特堡 feud.Barony = &克洛斯特堡Kloster_bergeBarony{}

func init() {
    f := BKloster_berge克洛斯特堡.(*克洛斯特堡Kloster_bergeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kloster_berge",
		TitleName: "克洛斯特堡",
		TitleCode: "b_kloster_berge",
	}
}
