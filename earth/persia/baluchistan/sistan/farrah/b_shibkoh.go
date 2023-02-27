package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 司阔霍ShibkohBarony struct {
	feud.BaseBarony
}

var BShibkoh司阔霍 feud.Barony = &司阔霍ShibkohBarony{}

func init() {
    f := BShibkoh司阔霍.(*司阔霍ShibkohBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shibkoh",
		TitleName: "司阔霍",
		TitleCode: "b_shibkoh",
	}
}
