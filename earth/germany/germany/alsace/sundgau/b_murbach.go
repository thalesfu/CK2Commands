package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔巴克MurbachBarony struct {
	feud.BaseBarony
}

var BMurbach米尔巴克 feud.Barony = &米尔巴克MurbachBarony{}

func init() {
    f := BMurbach米尔巴克.(*米尔巴克MurbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murbach",
		TitleName: "米尔巴克",
		TitleCode: "b_murbach",
	}
}
