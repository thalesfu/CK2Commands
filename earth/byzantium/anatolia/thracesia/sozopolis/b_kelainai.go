package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱奈KelainaiBarony struct {
	feud.BaseBarony
}

var BKelainai克莱奈 feud.Barony = &克莱奈KelainaiBarony{}

func init() {
    f := BKelainai克莱奈.(*克莱奈KelainaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelainai",
		TitleName: "克莱奈",
		TitleCode: "b_kelainai",
	}
}
