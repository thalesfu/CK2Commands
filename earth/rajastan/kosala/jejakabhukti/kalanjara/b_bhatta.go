package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋多BhattaBarony struct {
	feud.BaseBarony
}

var BBhatta跋多 feud.Barony = &跋多BhattaBarony{}

func init() {
	f := BBhatta跋多.(*跋多BhattaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhatta",
		TitleName: "跋多",
		TitleCode: "b_bhatta",
	}
}
