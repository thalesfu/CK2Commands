package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德日DerriBarony struct {
	feud.BaseBarony
}

var BDerri德日 feud.Barony = &德日DerriBarony{}

func init() {
    f := BDerri德日.(*德日DerriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derri",
		TitleName: "德日",
		TitleCode: "b_derri",
	}
}
