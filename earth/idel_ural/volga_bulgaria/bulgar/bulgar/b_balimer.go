package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒梅尔BalimerBarony struct {
	feud.BaseBarony
}

var BBalimer巴勒梅尔 feud.Barony = &巴勒梅尔BalimerBarony{}

func init() {
    f := BBalimer巴勒梅尔.(*巴勒梅尔BalimerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balimer",
		TitleName: "巴勒梅尔",
		TitleCode: "b_balimer",
	}
}
