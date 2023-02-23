package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔多BordeauxBarony struct {
	feud.BaseBarony
}

var BBordeaux波尔多 feud.Barony = &波尔多BordeauxBarony{}

func init() {
	f := BBordeaux波尔多.(*波尔多BordeauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bordeaux",
		TitleName: "波尔多",
		TitleCode: "b_bordeaux",
	}
}
