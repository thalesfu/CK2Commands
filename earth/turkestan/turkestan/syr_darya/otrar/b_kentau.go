package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯套KentauBarony struct {
	feud.BaseBarony
}

var BKentau肯套 feud.Barony = &肯套KentauBarony{}

func init() {
    f := BKentau肯套.(*肯套KentauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kentau",
		TitleName: "肯套",
		TitleCode: "b_kentau",
	}
}
