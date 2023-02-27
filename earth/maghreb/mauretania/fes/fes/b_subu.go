package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏布SubuBarony struct {
	feud.BaseBarony
}

var BSubu苏布 feud.Barony = &苏布SubuBarony{}

func init() {
    f := BSubu苏布.(*苏布SubuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subu",
		TitleName: "苏布",
		TitleCode: "b_subu",
	}
}
