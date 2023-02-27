package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯松KaisunBarony struct {
	feud.BaseBarony
}

var BKaisun凯松 feud.Barony = &凯松KaisunBarony{}

func init() {
    f := BKaisun凯松.(*凯松KaisunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaisun",
		TitleName: "凯松",
		TitleCode: "b_kaisun",
	}
}
