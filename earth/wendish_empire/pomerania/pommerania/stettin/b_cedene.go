package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采代内CedeneBarony struct {
	feud.BaseBarony
}

var BCedene采代内 feud.Barony = &采代内CedeneBarony{}

func init() {
    f := BCedene采代内.(*采代内CedeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cedene",
		TitleName: "采代内",
		TitleCode: "b_cedene",
	}
}
