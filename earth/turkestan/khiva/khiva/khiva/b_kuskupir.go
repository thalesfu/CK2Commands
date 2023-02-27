package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯库皮尔KuskupirBarony struct {
	feud.BaseBarony
}

var BKuskupir库斯库皮尔 feud.Barony = &库斯库皮尔KuskupirBarony{}

func init() {
    f := BKuskupir库斯库皮尔.(*库斯库皮尔KuskupirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuskupir",
		TitleName: "库斯库皮尔",
		TitleCode: "b_kuskupir",
	}
}
