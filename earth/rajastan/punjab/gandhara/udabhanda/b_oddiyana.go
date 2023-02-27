package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌苌OddiyanaBarony struct {
	feud.BaseBarony
}

var BOddiyana乌苌 feud.Barony = &乌苌OddiyanaBarony{}

func init() {
    f := BOddiyana乌苌.(*乌苌OddiyanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oddiyana",
		TitleName: "乌苌",
		TitleCode: "b_oddiyana",
	}
}
