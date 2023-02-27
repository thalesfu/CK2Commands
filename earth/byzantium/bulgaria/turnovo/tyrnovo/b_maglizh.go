package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默格利日MaglizhBarony struct {
	feud.BaseBarony
}

var BMaglizh默格利日 feud.Barony = &默格利日MaglizhBarony{}

func init() {
    f := BMaglizh默格利日.(*默格利日MaglizhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maglizh",
		TitleName: "默格利日",
		TitleCode: "b_maglizh",
	}
}
