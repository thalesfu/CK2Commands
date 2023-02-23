package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明屋MingoiBarony struct {
	feud.BaseBarony
}

var BMingoi明屋 feud.Barony = &明屋MingoiBarony{}

func init() {
	f := BMingoi明屋.(*明屋MingoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mingoi",
		TitleName: "明屋",
		TitleCode: "b_mingoi",
	}
}
