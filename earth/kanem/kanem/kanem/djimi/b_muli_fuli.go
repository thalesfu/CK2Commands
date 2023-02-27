package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆利富利Muli_fuliBarony struct {
	feud.BaseBarony
}

var BMuli_fuli穆利富利 feud.Barony = &穆利富利Muli_fuliBarony{}

func init() {
    f := BMuli_fuli穆利富利.(*穆利富利Muli_fuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muli_fuli",
		TitleName: "穆利富利",
		TitleCode: "b_muli_fuli",
	}
}
