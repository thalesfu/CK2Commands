package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多博伊DobojBarony struct {
	feud.BaseBarony
}

var BDoboj多博伊 feud.Barony = &多博伊DobojBarony{}

func init() {
	f := BDoboj多博伊.(*多博伊DobojBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doboj",
		TitleName: "多博伊",
		TitleCode: "b_doboj",
	}
}
