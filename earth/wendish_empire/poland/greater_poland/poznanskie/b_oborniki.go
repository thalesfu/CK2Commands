package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博尔尼基ObornikiBarony struct {
	feud.BaseBarony
}

var BOborniki奥博尔尼基 feud.Barony = &奥博尔尼基ObornikiBarony{}

func init() {
    f := BOborniki奥博尔尼基.(*奥博尔尼基ObornikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oborniki",
		TitleName: "奥博尔尼基",
		TitleCode: "b_oborniki",
	}
}
