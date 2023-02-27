package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴勒迪克BarleducBarony struct {
	feud.BaseBarony
}

var BBarleduc巴勒迪克 feud.Barony = &巴勒迪克BarleducBarony{}

func init() {
    f := BBarleduc巴勒迪克.(*巴勒迪克BarleducBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barleduc",
		TitleName: "巴勒迪克",
		TitleCode: "b_barleduc",
	}
}
