package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎博勒ZabolBarony struct {
	feud.BaseBarony
}

var BZabol扎博勒 feud.Barony = &扎博勒ZabolBarony{}

func init() {
    f := BZabol扎博勒.(*扎博勒ZabolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zabol",
		TitleName: "扎博勒",
		TitleCode: "b_zabol",
	}
}
