package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科拉莱斯CorralesBarony struct {
	feud.BaseBarony
}

var BCorrales科拉莱斯 feud.Barony = &科拉莱斯CorralesBarony{}

func init() {
    f := BCorrales科拉莱斯.(*科拉莱斯CorralesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corrales",
		TitleName: "科拉莱斯",
		TitleCode: "b_corrales",
	}
}
