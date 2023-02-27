package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博克西特BoksitiBarony struct {
	feud.BaseBarony
}

var BBoksiti博克西特 feud.Barony = &博克西特BoksitiBarony{}

func init() {
    f := BBoksiti博克西特.(*博克西特BoksitiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boksiti",
		TitleName: "博克西特",
		TitleCode: "b_boksiti",
	}
}
