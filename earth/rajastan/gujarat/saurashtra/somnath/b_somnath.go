package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏摩那他SomnathBarony struct {
	feud.BaseBarony
}

var BSomnath苏摩那他 feud.Barony = &苏摩那他SomnathBarony{}

func init() {
    f := BSomnath苏摩那他.(*苏摩那他SomnathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somnath",
		TitleName: "苏摩那他",
		TitleCode: "b_somnath",
	}
}
