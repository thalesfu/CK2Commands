package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁西斯CrusisBarony struct {
	feud.BaseBarony
}

var BCrusis克鲁西斯 feud.Barony = &克鲁西斯CrusisBarony{}

func init() {
    f := BCrusis克鲁西斯.(*克鲁西斯CrusisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crusis",
		TitleName: "克鲁西斯",
		TitleCode: "b_crusis",
	}
}
