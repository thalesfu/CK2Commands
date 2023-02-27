package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷彭施泰特ReppenstedtBarony struct {
	feud.BaseBarony
}

var BReppenstedt雷彭施泰特 feud.Barony = &雷彭施泰特ReppenstedtBarony{}

func init() {
    f := BReppenstedt雷彭施泰特.(*雷彭施泰特ReppenstedtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reppenstedt",
		TitleName: "雷彭施泰特",
		TitleCode: "b_reppenstedt",
	}
}
