package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥冈加OgangaBarony struct {
	feud.BaseBarony
}

var BOganga奥冈加 feud.Barony = &奥冈加OgangaBarony{}

func init() {
	f := BOganga奥冈加.(*奥冈加OgangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oganga",
		TitleName: "奥冈加",
		TitleCode: "b_oganga",
	}
}
