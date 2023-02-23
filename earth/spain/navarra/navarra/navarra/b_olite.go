package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利特OliteBarony struct {
	feud.BaseBarony
}

var BOlite奥利特 feud.Barony = &奥利特OliteBarony{}

func init() {
	f := BOlite奥利特.(*奥利特OliteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olite",
		TitleName: "奥利特",
		TitleCode: "b_olite",
	}
}
