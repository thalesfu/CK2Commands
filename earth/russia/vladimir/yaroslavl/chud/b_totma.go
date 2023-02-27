package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托季马TotmaBarony struct {
	feud.BaseBarony
}

var BTotma托季马 feud.Barony = &托季马TotmaBarony{}

func init() {
    f := BTotma托季马.(*托季马TotmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "totma",
		TitleName: "托季马",
		TitleCode: "b_totma",
	}
}
