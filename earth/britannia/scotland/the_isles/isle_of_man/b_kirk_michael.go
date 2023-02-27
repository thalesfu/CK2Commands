package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯克迈克尔Kirk_michaelBarony struct {
	feud.BaseBarony
}

var BKirk_michael柯克迈克尔 feud.Barony = &柯克迈克尔Kirk_michaelBarony{}

func init() {
    f := BKirk_michael柯克迈克尔.(*柯克迈克尔Kirk_michaelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirk_michael",
		TitleName: "柯克迈克尔",
		TitleCode: "b_kirk_michael",
	}
}
