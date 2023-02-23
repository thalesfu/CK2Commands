package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔克SrukBarony struct {
	feud.BaseBarony
}

var BSruk苏尔克 feud.Barony = &苏尔克SrukBarony{}

func init() {
	f := BSruk苏尔克.(*苏尔克SrukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sruk",
		TitleName: "苏尔克",
		TitleCode: "b_sruk",
	}
}
