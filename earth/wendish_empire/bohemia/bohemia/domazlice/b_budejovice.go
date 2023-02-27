package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布杰约维采BudejoviceBarony struct {
	feud.BaseBarony
}

var BBudejovice布杰约维采 feud.Barony = &布杰约维采BudejoviceBarony{}

func init() {
    f := BBudejovice布杰约维采.(*布杰约维采BudejoviceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "budejovice",
		TitleName: "布杰约维采",
		TitleCode: "b_budejovice",
	}
}
