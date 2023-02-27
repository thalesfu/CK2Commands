package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孟菲斯MemphisBarony struct {
	feud.BaseBarony
}

var BMemphis孟菲斯 feud.Barony = &孟菲斯MemphisBarony{}

func init() {
    f := BMemphis孟菲斯.(*孟菲斯MemphisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "memphis",
		TitleName: "孟菲斯",
		TitleCode: "b_memphis",
	}
}
