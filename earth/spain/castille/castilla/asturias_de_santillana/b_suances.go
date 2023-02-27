package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏安塞斯SuancesBarony struct {
	feud.BaseBarony
}

var BSuances苏安塞斯 feud.Barony = &苏安塞斯SuancesBarony{}

func init() {
    f := BSuances苏安塞斯.(*苏安塞斯SuancesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suances",
		TitleName: "苏安塞斯",
		TitleCode: "b_suances",
	}
}
