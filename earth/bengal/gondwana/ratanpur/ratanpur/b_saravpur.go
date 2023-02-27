package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑罗婆城SaravpurBarony struct {
	feud.BaseBarony
}

var BSaravpur娑罗婆城 feud.Barony = &娑罗婆城SaravpurBarony{}

func init() {
    f := BSaravpur娑罗婆城.(*娑罗婆城SaravpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saravpur",
		TitleName: "娑罗婆城",
		TitleCode: "b_saravpur",
	}
}
