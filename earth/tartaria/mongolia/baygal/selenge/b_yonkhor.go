package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 永霍尔YonkhorBarony struct {
	feud.BaseBarony
}

var BYonkhor永霍尔 feud.Barony = &永霍尔YonkhorBarony{}

func init() {
    f := BYonkhor永霍尔.(*永霍尔YonkhorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yonkhor",
		TitleName: "永霍尔",
		TitleCode: "b_yonkhor",
	}
}
