package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯托加AstorgaBarony struct {
	feud.BaseBarony
}

var BAstorga阿斯托加 feud.Barony = &阿斯托加AstorgaBarony{}

func init() {
	f := BAstorga阿斯托加.(*阿斯托加AstorgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "astorga",
		TitleName: "阿斯托加",
		TitleCode: "b_astorga",
	}
}
