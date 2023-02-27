package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济班DhibanBarony struct {
	feud.BaseBarony
}

var BDhiban济班 feud.Barony = &济班DhibanBarony{}

func init() {
    f := BDhiban济班.(*济班DhibanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhiban",
		TitleName: "济班",
		TitleCode: "b_dhiban",
	}
}
