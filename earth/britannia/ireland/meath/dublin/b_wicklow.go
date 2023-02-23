package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威克洛WicklowBarony struct {
	feud.BaseBarony
}

var BWicklow威克洛 feud.Barony = &威克洛WicklowBarony{}

func init() {
	f := BWicklow威克洛.(*威克洛WicklowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wicklow",
		TitleName: "威克洛",
		TitleCode: "b_wicklow",
	}
}
