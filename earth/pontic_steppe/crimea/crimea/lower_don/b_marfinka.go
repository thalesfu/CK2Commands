package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔芬卡MarfinkaBarony struct {
	feud.BaseBarony
}

var BMarfinka马尔芬卡 feud.Barony = &马尔芬卡MarfinkaBarony{}

func init() {
    f := BMarfinka马尔芬卡.(*马尔芬卡MarfinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marfinka",
		TitleName: "马尔芬卡",
		TitleCode: "b_marfinka",
	}
}
