package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加巴苏尔GabasourBarony struct {
	feud.BaseBarony
}

var BGabasour加巴苏尔 feud.Barony = &加巴苏尔GabasourBarony{}

func init() {
    f := BGabasour加巴苏尔.(*加巴苏尔GabasourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabasour",
		TitleName: "加巴苏尔",
		TitleCode: "b_gabasour",
	}
}
