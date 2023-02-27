package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥特里科利OtricoliBarony struct {
	feud.BaseBarony
}

var BOtricoli奥特里科利 feud.Barony = &奥特里科利OtricoliBarony{}

func init() {
    f := BOtricoli奥特里科利.(*奥特里科利OtricoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otricoli",
		TitleName: "奥特里科利",
		TitleCode: "b_otricoli",
	}
}
