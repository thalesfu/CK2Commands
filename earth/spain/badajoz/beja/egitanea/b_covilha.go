package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科维良CovilhaBarony struct {
	feud.BaseBarony
}

var BCovilha科维良 feud.Barony = &科维良CovilhaBarony{}

func init() {
    f := BCovilha科维良.(*科维良CovilhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "covilha",
		TitleName: "科维良",
		TitleCode: "b_covilha",
	}
}
