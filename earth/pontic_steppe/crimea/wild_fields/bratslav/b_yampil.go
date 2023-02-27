package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扬皮尔YampilBarony struct {
	feud.BaseBarony
}

var BYampil扬皮尔 feud.Barony = &扬皮尔YampilBarony{}

func init() {
    f := BYampil扬皮尔.(*扬皮尔YampilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yampil",
		TitleName: "扬皮尔",
		TitleCode: "b_yampil",
	}
}
