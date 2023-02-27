package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肯吉尔KengirBarony struct {
	feud.BaseBarony
}

var BKengir肯吉尔 feud.Barony = &肯吉尔KengirBarony{}

func init() {
    f := BKengir肯吉尔.(*肯吉尔KengirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kengir",
		TitleName: "肯吉尔",
		TitleCode: "b_kengir",
	}
}
