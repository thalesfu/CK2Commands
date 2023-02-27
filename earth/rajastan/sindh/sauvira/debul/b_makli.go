package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩柯尼MakliBarony struct {
	feud.BaseBarony
}

var BMakli摩柯尼 feud.Barony = &摩柯尼MakliBarony{}

func init() {
    f := BMakli摩柯尼.(*摩柯尼MakliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makli",
		TitleName: "摩柯尼",
		TitleCode: "b_makli",
	}
}
