package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔塔涅MortagnedevosgesBarony struct {
	feud.BaseBarony
}

var BMortagnedevosges莫尔塔涅 feud.Barony = &莫尔塔涅MortagnedevosgesBarony{}

func init() {
    f := BMortagnedevosges莫尔塔涅.(*莫尔塔涅MortagnedevosgesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mortagnedevosges",
		TitleName: "莫尔塔涅",
		TitleCode: "b_mortagnedevosges",
	}
}
