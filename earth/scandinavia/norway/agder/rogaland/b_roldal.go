package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒尔达尔RoldalBarony struct {
	feud.BaseBarony
}

var BRoldal勒尔达尔 feud.Barony = &勒尔达尔RoldalBarony{}

func init() {
    f := BRoldal勒尔达尔.(*勒尔达尔RoldalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roldal",
		TitleName: "勒尔达尔",
		TitleCode: "b_roldal",
	}
}
