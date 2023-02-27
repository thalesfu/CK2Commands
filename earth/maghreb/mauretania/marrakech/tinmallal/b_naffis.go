package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳菲斯NaffisBarony struct {
	feud.BaseBarony
}

var BNaffis纳菲斯 feud.Barony = &纳菲斯NaffisBarony{}

func init() {
    f := BNaffis纳菲斯.(*纳菲斯NaffisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naffis",
		TitleName: "纳菲斯",
		TitleCode: "b_naffis",
	}
}
