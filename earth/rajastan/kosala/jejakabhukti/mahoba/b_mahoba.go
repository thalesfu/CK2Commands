package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩呼婆MahobaBarony struct {
	feud.BaseBarony
}

var BMahoba摩呼婆 feud.Barony = &摩呼婆MahobaBarony{}

func init() {
    f := BMahoba摩呼婆.(*摩呼婆MahobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahoba",
		TitleName: "摩呼婆",
		TitleCode: "b_mahoba",
	}
}
