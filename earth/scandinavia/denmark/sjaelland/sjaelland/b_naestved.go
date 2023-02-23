package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈斯特韦兹NaestvedBarony struct {
	feud.BaseBarony
}

var BNaestved奈斯特韦兹 feud.Barony = &奈斯特韦兹NaestvedBarony{}

func init() {
	f := BNaestved奈斯特韦兹.(*奈斯特韦兹NaestvedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naestved",
		TitleName: "奈斯特韦兹",
		TitleCode: "b_naestved",
	}
}
