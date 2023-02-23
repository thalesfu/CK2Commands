package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎塔涅迪CantanhedeBarony struct {
	feud.BaseBarony
}

var BCantanhede坎塔涅迪 feud.Barony = &坎塔涅迪CantanhedeBarony{}

func init() {
	f := BCantanhede坎塔涅迪.(*坎塔涅迪CantanhedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cantanhede",
		TitleName: "坎塔涅迪",
		TitleCode: "b_cantanhede",
	}
}
