package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈马MaymaBarony struct {
	feud.BaseBarony
}

var BMayma迈马 feud.Barony = &迈马MaymaBarony{}

func init() {
    f := BMayma迈马.(*迈马MaymaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayma",
		TitleName: "迈马",
		TitleCode: "b_mayma",
	}
}
