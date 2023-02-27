package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博戈柳博沃BogolyubovoBarony struct {
	feud.BaseBarony
}

var BBogolyubovo博戈柳博沃 feud.Barony = &博戈柳博沃BogolyubovoBarony{}

func init() {
    f := BBogolyubovo博戈柳博沃.(*博戈柳博沃BogolyubovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogolyubovo",
		TitleName: "博戈柳博沃",
		TitleCode: "b_bogolyubovo",
	}
}
