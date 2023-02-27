package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣莫洛格St_moluagBarony struct {
	feud.BaseBarony
}

var BSt_moluag圣莫洛格 feud.Barony = &圣莫洛格St_moluagBarony{}

func init() {
    f := BSt_moluag圣莫洛格.(*圣莫洛格St_moluagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_moluag",
		TitleName: "圣莫洛格",
		TitleCode: "b_st_moluag",
	}
}
