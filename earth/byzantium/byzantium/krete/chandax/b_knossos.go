package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克诺索斯KnossosBarony struct {
	feud.BaseBarony
}

var BKnossos克诺索斯 feud.Barony = &克诺索斯KnossosBarony{}

func init() {
    f := BKnossos克诺索斯.(*克诺索斯KnossosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knossos",
		TitleName: "克诺索斯",
		TitleCode: "b_knossos",
	}
}
