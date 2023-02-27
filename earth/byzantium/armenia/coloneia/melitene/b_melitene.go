package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅利蒂尼MeliteneBarony struct {
	feud.BaseBarony
}

var BMelitene梅利蒂尼 feud.Barony = &梅利蒂尼MeliteneBarony{}

func init() {
    f := BMelitene梅利蒂尼.(*梅利蒂尼MeliteneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melitene",
		TitleName: "梅利蒂尼",
		TitleCode: "b_melitene",
	}
}
