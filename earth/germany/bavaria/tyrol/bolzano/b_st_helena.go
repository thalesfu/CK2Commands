package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣海伦娜St_helenaBarony struct {
	feud.BaseBarony
}

var BSt_helena圣海伦娜 feud.Barony = &圣海伦娜St_helenaBarony{}

func init() {
    f := BSt_helena圣海伦娜.(*圣海伦娜St_helenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_helena",
		TitleName: "圣海伦娜",
		TitleCode: "b_st_helena",
	}
}
