package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿丹St_aidansBarony struct {
	feud.BaseBarony
}

var BSt_aidans圣阿丹 feud.Barony = &圣阿丹St_aidansBarony{}

func init() {
    f := BSt_aidans圣阿丹.(*圣阿丹St_aidansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_aidans",
		TitleName: "圣阿丹",
		TitleCode: "b_st_aidans",
	}
}
