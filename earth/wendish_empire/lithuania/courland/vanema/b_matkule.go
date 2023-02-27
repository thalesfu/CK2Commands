package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马特库莱MatkuleBarony struct {
	feud.BaseBarony
}

var BMatkule马特库莱 feud.Barony = &马特库莱MatkuleBarony{}

func init() {
    f := BMatkule马特库莱.(*马特库莱MatkuleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matkule",
		TitleName: "马特库莱",
		TitleCode: "b_matkule",
	}
}
