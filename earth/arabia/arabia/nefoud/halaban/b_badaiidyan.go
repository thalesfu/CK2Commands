package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴达鲁杜安BadaiidyanBarony struct {
	feud.BaseBarony
}

var BBadaiidyan巴达鲁杜安 feud.Barony = &巴达鲁杜安BadaiidyanBarony{}

func init() {
    f := BBadaiidyan巴达鲁杜安.(*巴达鲁杜安BadaiidyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badaiidyan",
		TitleName: "巴达鲁杜安",
		TitleCode: "b_badaiidyan",
	}
}
