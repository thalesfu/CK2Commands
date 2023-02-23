package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 骑士堡KrakdeschevaliersBarony struct {
	feud.BaseBarony
}

var BKrakdeschevaliers骑士堡 feud.Barony = &骑士堡KrakdeschevaliersBarony{}

func init() {
	f := BKrakdeschevaliers骑士堡.(*骑士堡KrakdeschevaliersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krakdeschevaliers",
		TitleName: "骑士堡",
		TitleCode: "b_krakdeschevaliers",
	}
}
