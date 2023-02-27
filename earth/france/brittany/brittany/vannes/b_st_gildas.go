package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣吉尔达St_gildasBarony struct {
	feud.BaseBarony
}

var BSt_gildas圣吉尔达 feud.Barony = &圣吉尔达St_gildasBarony{}

func init() {
    f := BSt_gildas圣吉尔达.(*圣吉尔达St_gildasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_gildas",
		TitleName: "圣吉尔达",
		TitleCode: "b_st_gildas",
	}
}
