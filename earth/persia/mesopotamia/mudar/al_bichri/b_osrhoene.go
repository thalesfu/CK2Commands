package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯霍尼OsrhoeneBarony struct {
	feud.BaseBarony
}

var BOsrhoene奥斯霍尼 feud.Barony = &奥斯霍尼OsrhoeneBarony{}

func init() {
    f := BOsrhoene奥斯霍尼.(*奥斯霍尼OsrhoeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osrhoene",
		TitleName: "奥斯霍尼",
		TitleCode: "b_osrhoene",
	}
}
