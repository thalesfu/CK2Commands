package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯摩伐底HemavatiBarony struct {
	feud.BaseBarony
}

var BHemavati醯摩伐底 feud.Barony = &醯摩伐底HemavatiBarony{}

func init() {
    f := BHemavati醯摩伐底.(*醯摩伐底HemavatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hemavati",
		TitleName: "醯摩伐底",
		TitleCode: "b_hemavati",
	}
}
