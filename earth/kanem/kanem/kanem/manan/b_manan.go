package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马南MananBarony struct {
	feud.BaseBarony
}

var BManan马南 feud.Barony = &马南MananBarony{}

func init() {
    f := BManan马南.(*马南MananBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manan",
		TitleName: "马南",
		TitleCode: "b_manan",
	}
}
