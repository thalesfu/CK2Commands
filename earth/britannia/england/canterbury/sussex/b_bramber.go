package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰伯BramberBarony struct {
	feud.BaseBarony
}

var BBramber布兰伯 feud.Barony = &布兰伯BramberBarony{}

func init() {
	f := BBramber布兰伯.(*布兰伯BramberBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bramber",
		TitleName: "布兰伯",
		TitleCode: "b_bramber",
	}
}
