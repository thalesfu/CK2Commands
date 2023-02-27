package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯陆OsloBarony struct {
	feud.BaseBarony
}

var BOslo奥斯陆 feud.Barony = &奥斯陆OsloBarony{}

func init() {
    f := BOslo奥斯陆.(*奥斯陆OsloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oslo",
		TitleName: "奥斯陆",
		TitleCode: "b_oslo",
	}
}
