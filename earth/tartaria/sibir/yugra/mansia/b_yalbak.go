package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尔巴克YalbakBarony struct {
	feud.BaseBarony
}

var BYalbak亚尔巴克 feud.Barony = &亚尔巴克YalbakBarony{}

func init() {
	f := BYalbak亚尔巴克.(*亚尔巴克YalbakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yalbak",
		TitleName: "亚尔巴克",
		TitleCode: "b_yalbak",
	}
}
