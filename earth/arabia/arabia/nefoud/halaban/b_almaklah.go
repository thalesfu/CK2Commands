package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈克拉AlmaklahBarony struct {
	feud.BaseBarony
}

var BAlmaklah迈克拉 feud.Barony = &迈克拉AlmaklahBarony{}

func init() {
    f := BAlmaklah迈克拉.(*迈克拉AlmaklahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almaklah",
		TitleName: "迈克拉",
		TitleCode: "b_almaklah",
	}
}
