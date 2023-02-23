package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴巴克BabakBarony struct {
	feud.BaseBarony
}

var BBabak巴巴克 feud.Barony = &巴巴克BabakBarony{}

func init() {
	f := BBabak巴巴克.(*巴巴克BabakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babak",
		TitleName: "巴巴克",
		TitleCode: "b_babak",
	}
}
