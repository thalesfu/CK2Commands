package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚济马VyazmaBarony struct {
	feud.BaseBarony
}

var BVyazma维亚济马 feud.Barony = &维亚济马VyazmaBarony{}

func init() {
    f := BVyazma维亚济马.(*维亚济马VyazmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyazma",
		TitleName: "维亚济马",
		TitleCode: "b_vyazma",
	}
}
