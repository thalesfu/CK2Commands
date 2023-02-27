package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣于尔萨讷St_ursanneBarony struct {
	feud.BaseBarony
}

var BSt_ursanne圣于尔萨讷 feud.Barony = &圣于尔萨讷St_ursanneBarony{}

func init() {
    f := BSt_ursanne圣于尔萨讷.(*圣于尔萨讷St_ursanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_ursanne",
		TitleName: "圣于尔萨讷",
		TitleCode: "b_st_ursanne",
	}
}
