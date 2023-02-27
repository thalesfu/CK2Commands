package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉苏泰赖讷LasouterraineBarony struct {
	feud.BaseBarony
}

var BLasouterraine拉苏泰赖讷 feud.Barony = &拉苏泰赖讷LasouterraineBarony{}

func init() {
    f := BLasouterraine拉苏泰赖讷.(*拉苏泰赖讷LasouterraineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lasouterraine",
		TitleName: "拉苏泰赖讷",
		TitleCode: "b_lasouterraine",
	}
}
