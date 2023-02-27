package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德BardBarony struct {
	feud.BaseBarony
}

var BBard巴德 feud.Barony = &巴德BardBarony{}

func init() {
    f := BBard巴德.(*巴德BardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bard",
		TitleName: "巴德",
		TitleCode: "b_bard",
	}
}
