package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣迈克尔St_michaelBarony struct {
	feud.BaseBarony
}

var BSt_michael圣迈克尔 feud.Barony = &圣迈克尔St_michaelBarony{}

func init() {
    f := BSt_michael圣迈克尔.(*圣迈克尔St_michaelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_michael",
		TitleName: "圣迈克尔",
		TitleCode: "b_st_michael",
	}
}
