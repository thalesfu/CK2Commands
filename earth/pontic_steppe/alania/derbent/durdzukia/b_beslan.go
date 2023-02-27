package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别斯兰BeslanBarony struct {
	feud.BaseBarony
}

var BBeslan别斯兰 feud.Barony = &别斯兰BeslanBarony{}

func init() {
    f := BBeslan别斯兰.(*别斯兰BeslanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beslan",
		TitleName: "别斯兰",
		TitleCode: "b_beslan",
	}
}
