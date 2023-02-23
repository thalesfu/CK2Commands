package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔达尼亚CerdanaBarony struct {
	feud.BaseBarony
}

var BCerdana塞尔达尼亚 feud.Barony = &塞尔达尼亚CerdanaBarony{}

func init() {
	f := BCerdana塞尔达尼亚.(*塞尔达尼亚CerdanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cerdana",
		TitleName: "塞尔达尼亚",
		TitleCode: "b_cerdana",
	}
}
