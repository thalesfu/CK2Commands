package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陈豆利DindoriBarony struct {
	feud.BaseBarony
}

var BDindori陈豆利 feud.Barony = &陈豆利DindoriBarony{}

func init() {
    f := BDindori陈豆利.(*陈豆利DindoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dindori",
		TitleName: "陈豆利",
		TitleCode: "b_dindori",
	}
}
