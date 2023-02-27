package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米什科尔茨MiskolcBarony struct {
	feud.BaseBarony
}

var BMiskolc米什科尔茨 feud.Barony = &米什科尔茨MiskolcBarony{}

func init() {
    f := BMiskolc米什科尔茨.(*米什科尔茨MiskolcBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miskolc",
		TitleName: "米什科尔茨",
		TitleCode: "b_miskolc",
	}
}
