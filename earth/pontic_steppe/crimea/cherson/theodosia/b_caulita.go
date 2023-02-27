package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考里塔CaulitaBarony struct {
	feud.BaseBarony
}

var BCaulita考里塔 feud.Barony = &考里塔CaulitaBarony{}

func init() {
    f := BCaulita考里塔.(*考里塔CaulitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caulita",
		TitleName: "考里塔",
		TitleCode: "b_caulita",
	}
}
