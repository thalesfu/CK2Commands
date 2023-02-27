package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳恩堡DanlauenburgBarony struct {
	feud.BaseBarony
}

var BDanlauenburg劳恩堡 feud.Barony = &劳恩堡DanlauenburgBarony{}

func init() {
    f := BDanlauenburg劳恩堡.(*劳恩堡DanlauenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danlauenburg",
		TitleName: "劳恩堡",
		TitleCode: "b_danlauenburg",
	}
}
