package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉讷GuinesBarony struct {
	feud.BaseBarony
}

var BGuines吉讷 feud.Barony = &吉讷GuinesBarony{}

func init() {
    f := BGuines吉讷.(*吉讷GuinesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guines",
		TitleName: "吉讷",
		TitleCode: "b_guines",
	}
}
