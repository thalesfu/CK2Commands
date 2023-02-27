package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲尔凯尔Fir_cellBarony struct {
	feud.BaseBarony
}

var BFir_cell菲尔凯尔 feud.Barony = &菲尔凯尔Fir_cellBarony{}

func init() {
    f := BFir_cell菲尔凯尔.(*菲尔凯尔Fir_cellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fir_cell",
		TitleName: "菲尔凯尔",
		TitleCode: "b_fir_cell",
	}
}
