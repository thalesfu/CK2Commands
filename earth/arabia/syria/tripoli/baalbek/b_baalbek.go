package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒贝克BaalbekBarony struct {
	feud.BaseBarony
}

var BBaalbek巴勒贝克 feud.Barony = &巴勒贝克BaalbekBarony{}

func init() {
    f := BBaalbek巴勒贝克.(*巴勒贝克BaalbekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baalbek",
		TitleName: "巴勒贝克",
		TitleCode: "b_baalbek",
	}
}
