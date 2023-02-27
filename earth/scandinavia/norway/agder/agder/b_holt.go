package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔特HoltBarony struct {
	feud.BaseBarony
}

var BHolt霍尔特 feud.Barony = &霍尔特HoltBarony{}

func init() {
    f := BHolt霍尔特.(*霍尔特HoltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holt",
		TitleName: "霍尔特",
		TitleCode: "b_holt",
	}
}
