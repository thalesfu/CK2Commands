package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯拉韦德GislavedBarony struct {
	feud.BaseBarony
}

var BGislaved伊斯拉韦德 feud.Barony = &伊斯拉韦德GislavedBarony{}

func init() {
    f := BGislaved伊斯拉韦德.(*伊斯拉韦德GislavedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gislaved",
		TitleName: "伊斯拉韦德",
		TitleCode: "b_gislaved",
	}
}
