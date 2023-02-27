package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗马黑尔DromahairBarony struct {
	feud.BaseBarony
}

var BDromahair德罗马黑尔 feud.Barony = &德罗马黑尔DromahairBarony{}

func init() {
    f := BDromahair德罗马黑尔.(*德罗马黑尔DromahairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dromahair",
		TitleName: "德罗马黑尔",
		TitleCode: "b_dromahair",
	}
}
