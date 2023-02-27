package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯BathBarony struct {
	feud.BaseBarony
}

var BBath巴斯 feud.Barony = &巴斯BathBarony{}

func init() {
    f := BBath巴斯.(*巴斯BathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bath",
		TitleName: "巴斯",
		TitleCode: "b_bath",
	}
}
