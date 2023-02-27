package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下奇尔NizhchirBarony struct {
	feud.BaseBarony
}

var BNizhchir下奇尔 feud.Barony = &下奇尔NizhchirBarony{}

func init() {
    f := BNizhchir下奇尔.(*下奇尔NizhchirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhchir",
		TitleName: "下奇尔",
		TitleCode: "b_nizhchir",
	}
}
