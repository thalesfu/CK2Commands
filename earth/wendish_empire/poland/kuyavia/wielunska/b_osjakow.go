package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥夏库夫OsjakowBarony struct {
	feud.BaseBarony
}

var BOsjakow奥夏库夫 feud.Barony = &奥夏库夫OsjakowBarony{}

func init() {
    f := BOsjakow奥夏库夫.(*奥夏库夫OsjakowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osjakow",
		TitleName: "奥夏库夫",
		TitleCode: "b_osjakow",
	}
}
