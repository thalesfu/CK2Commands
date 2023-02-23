package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 的黎波里LibtripoliBarony struct {
	feud.BaseBarony
}

var BLibtripoli的黎波里 feud.Barony = &的黎波里LibtripoliBarony{}

func init() {
	f := BLibtripoli的黎波里.(*的黎波里LibtripoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "libtripoli",
		TitleName: "的黎波里",
		TitleCode: "b_libtripoli",
	}
}
