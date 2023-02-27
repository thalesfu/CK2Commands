package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆卢利BaroliBarony struct {
	feud.BaseBarony
}

var BBaroli婆卢利 feud.Barony = &婆卢利BaroliBarony{}

func init() {
    f := BBaroli婆卢利.(*婆卢利BaroliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baroli",
		TitleName: "婆卢利",
		TitleCode: "b_baroli",
	}
}
