package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴日马克BajmaqBarony struct {
	feud.BaseBarony
}

var BBajmaq巴日马克 feud.Barony = &巴日马克BajmaqBarony{}

func init() {
    f := BBajmaq巴日马克.(*巴日马克BajmaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bajmaq",
		TitleName: "巴日马克",
		TitleCode: "b_bajmaq",
	}
}
