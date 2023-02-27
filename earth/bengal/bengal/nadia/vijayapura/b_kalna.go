package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗那KalnaBarony struct {
	feud.BaseBarony
}

var BKalna迦罗那 feud.Barony = &迦罗那KalnaBarony{}

func init() {
    f := BKalna迦罗那.(*迦罗那KalnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalna",
		TitleName: "迦罗那",
		TitleCode: "b_kalna",
	}
}
