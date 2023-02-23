package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格布勒BagbleBarony struct {
	feud.BaseBarony
}

var BBagble巴格布勒 feud.Barony = &巴格布勒BagbleBarony{}

func init() {
	f := BBagble巴格布勒.(*巴格布勒BagbleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagble",
		TitleName: "巴格布勒",
		TitleCode: "b_bagble",
	}
}
