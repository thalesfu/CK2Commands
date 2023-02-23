package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马扎拉MazaraBarony struct {
	feud.BaseBarony
}

var BMazara马扎拉 feud.Barony = &马扎拉MazaraBarony{}

func init() {
	f := BMazara马扎拉.(*马扎拉MazaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazara",
		TitleName: "马扎拉",
		TitleCode: "b_mazara",
	}
}
