package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔内SarnyBarony struct {
	feud.BaseBarony
}

var BSarny萨尔内 feud.Barony = &萨尔内SarnyBarony{}

func init() {
    f := BSarny萨尔内.(*萨尔内SarnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarny",
		TitleName: "萨尔内",
		TitleCode: "b_sarny",
	}
}
