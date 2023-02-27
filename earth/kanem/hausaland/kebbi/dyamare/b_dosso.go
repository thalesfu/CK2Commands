package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多索DossoBarony struct {
	feud.BaseBarony
}

var BDosso多索 feud.Barony = &多索DossoBarony{}

func init() {
    f := BDosso多索.(*多索DossoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dosso",
		TitleName: "多索",
		TitleCode: "b_dosso",
	}
}
