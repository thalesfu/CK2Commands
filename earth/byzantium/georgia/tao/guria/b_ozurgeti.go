package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥祖尔盖蒂OzurgetiBarony struct {
	feud.BaseBarony
}

var BOzurgeti奥祖尔盖蒂 feud.Barony = &奥祖尔盖蒂OzurgetiBarony{}

func init() {
	f := BOzurgeti奥祖尔盖蒂.(*奥祖尔盖蒂OzurgetiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozurgeti",
		TitleName: "奥祖尔盖蒂",
		TitleCode: "b_ozurgeti",
	}
}
