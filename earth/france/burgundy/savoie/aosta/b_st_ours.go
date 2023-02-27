package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣乌尔St_oursBarony struct {
	feud.BaseBarony
}

var BSt_ours圣乌尔 feud.Barony = &圣乌尔St_oursBarony{}

func init() {
    f := BSt_ours圣乌尔.(*圣乌尔St_oursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_ours",
		TitleName: "圣乌尔",
		TitleCode: "b_st_ours",
	}
}
