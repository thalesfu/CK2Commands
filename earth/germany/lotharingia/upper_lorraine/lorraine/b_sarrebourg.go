package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔堡SarrebourgBarony struct {
	feud.BaseBarony
}

var BSarrebourg萨尔堡 feud.Barony = &萨尔堡SarrebourgBarony{}

func init() {
    f := BSarrebourg萨尔堡.(*萨尔堡SarrebourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarrebourg",
		TitleName: "萨尔堡",
		TitleCode: "b_sarrebourg",
	}
}
