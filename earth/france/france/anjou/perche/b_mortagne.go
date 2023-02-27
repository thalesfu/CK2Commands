package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔塔涅MortagneBarony struct {
	feud.BaseBarony
}

var BMortagne莫尔塔涅 feud.Barony = &莫尔塔涅MortagneBarony{}

func init() {
    f := BMortagne莫尔塔涅.(*莫尔塔涅MortagneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mortagne",
		TitleName: "莫尔塔涅",
		TitleCode: "b_mortagne",
	}
}
