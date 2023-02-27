package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔苏埃格拉ConsuegraBarony struct {
	feud.BaseBarony
}

var BConsuegra孔苏埃格拉 feud.Barony = &孔苏埃格拉ConsuegraBarony{}

func init() {
    f := BConsuegra孔苏埃格拉.(*孔苏埃格拉ConsuegraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "consuegra",
		TitleName: "孔苏埃格拉",
		TitleCode: "b_consuegra",
	}
}
