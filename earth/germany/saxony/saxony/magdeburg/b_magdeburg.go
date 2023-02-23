package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格德堡MagdeburgBarony struct {
	feud.BaseBarony
}

var BMagdeburg马格德堡 feud.Barony = &马格德堡MagdeburgBarony{}

func init() {
	f := BMagdeburg马格德堡.(*马格德堡MagdeburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magdeburg",
		TitleName: "马格德堡",
		TitleCode: "b_magdeburg",
	}
}
