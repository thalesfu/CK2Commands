package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪苏格DisuqBarony struct {
	feud.BaseBarony
}

var BDisuq迪苏格 feud.Barony = &迪苏格DisuqBarony{}

func init() {
	f := BDisuq迪苏格.(*迪苏格DisuqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "disuq",
		TitleName: "迪苏格",
		TitleCode: "b_disuq",
	}
}
