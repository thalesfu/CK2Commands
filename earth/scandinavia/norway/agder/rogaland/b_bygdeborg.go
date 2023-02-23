package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布格德堡BygdeborgBarony struct {
	feud.BaseBarony
}

var BBygdeborg布格德堡 feud.Barony = &布格德堡BygdeborgBarony{}

func init() {
	f := BBygdeborg布格德堡.(*布格德堡BygdeborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bygdeborg",
		TitleName: "布格德堡",
		TitleCode: "b_bygdeborg",
	}
}
