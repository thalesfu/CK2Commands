package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门士MonicerBarony struct {
	feud.BaseBarony
}

var BMonicer门士 feud.Barony = &门士MonicerBarony{}

func init() {
	f := BMonicer门士.(*门士MonicerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monicer",
		TitleName: "门士",
		TitleCode: "b_monicer",
	}
}
