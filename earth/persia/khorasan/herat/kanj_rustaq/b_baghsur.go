package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格苏尔BaghsurBarony struct {
	feud.BaseBarony
}

var BBaghsur巴格苏尔 feud.Barony = &巴格苏尔BaghsurBarony{}

func init() {
    f := BBaghsur巴格苏尔.(*巴格苏尔BaghsurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghsur",
		TitleName: "巴格苏尔",
		TitleCode: "b_baghsur",
	}
}
