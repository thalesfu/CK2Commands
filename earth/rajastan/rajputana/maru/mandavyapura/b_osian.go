package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌珊OsianBarony struct {
	feud.BaseBarony
}

var BOsian乌珊 feud.Barony = &乌珊OsianBarony{}

func init() {
	f := BOsian乌珊.(*乌珊OsianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osian",
		TitleName: "乌珊",
		TitleCode: "b_osian",
	}
}
