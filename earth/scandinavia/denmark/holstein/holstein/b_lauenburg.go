package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳恩堡LauenburgBarony struct {
	feud.BaseBarony
}

var BLauenburg劳恩堡 feud.Barony = &劳恩堡LauenburgBarony{}

func init() {
    f := BLauenburg劳恩堡.(*劳恩堡LauenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lauenburg",
		TitleName: "劳恩堡",
		TitleCode: "b_lauenburg",
	}
}
