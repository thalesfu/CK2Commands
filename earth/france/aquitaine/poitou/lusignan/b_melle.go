package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅勒MelleBarony struct {
	feud.BaseBarony
}

var BMelle梅勒 feud.Barony = &梅勒MelleBarony{}

func init() {
	f := BMelle梅勒.(*梅勒MelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melle",
		TitleName: "梅勒",
		TitleCode: "b_melle",
	}
}
