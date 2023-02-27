package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡米什利QamishhloBarony struct {
	feud.BaseBarony
}

var BQamishhlo卡米什利 feud.Barony = &卡米什利QamishhloBarony{}

func init() {
    f := BQamishhlo卡米什利.(*卡米什利QamishhloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qamishhlo",
		TitleName: "卡米什利",
		TitleCode: "b_qamishhlo",
	}
}
