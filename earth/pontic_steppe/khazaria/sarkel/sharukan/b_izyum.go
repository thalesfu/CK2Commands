package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊久姆IzyumBarony struct {
	feud.BaseBarony
}

var BIzyum伊久姆 feud.Barony = &伊久姆IzyumBarony{}

func init() {
    f := BIzyum伊久姆.(*伊久姆IzyumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izyum",
		TitleName: "伊久姆",
		TitleCode: "b_izyum",
	}
}
