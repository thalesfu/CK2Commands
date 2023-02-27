package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔马PalmaBarony struct {
	feud.BaseBarony
}

var BPalma帕尔马 feud.Barony = &帕尔马PalmaBarony{}

func init() {
    f := BPalma帕尔马.(*帕尔马PalmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palma",
		TitleName: "帕尔马",
		TitleCode: "b_palma",
	}
}
