package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴库BakuBarony struct {
	feud.BaseBarony
}

var BBaku巴库 feud.Barony = &巴库BakuBarony{}

func init() {
    f := BBaku巴库.(*巴库BakuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baku",
		TitleName: "巴库",
		TitleCode: "b_baku",
	}
}
