package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣保罗St_paulsBarony struct {
	feud.BaseBarony
}

var BSt_pauls圣保罗 feud.Barony = &圣保罗St_paulsBarony{}

func init() {
    f := BSt_pauls圣保罗.(*圣保罗St_paulsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_pauls",
		TitleName: "圣保罗",
		TitleCode: "b_st_pauls",
	}
}
