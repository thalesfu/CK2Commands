package mathura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勿邻陀林VrindavanBarony struct {
	feud.BaseBarony
}

var BVrindavan勿邻陀林 feud.Barony = &勿邻陀林VrindavanBarony{}

func init() {
	f := BVrindavan勿邻陀林.(*勿邻陀林VrindavanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vrindavan",
		TitleName: "勿邻陀林",
		TitleCode: "b_vrindavan",
	}
}
