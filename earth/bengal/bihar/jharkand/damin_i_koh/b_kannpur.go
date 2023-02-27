package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乾补KannpurBarony struct {
	feud.BaseBarony
}

var BKannpur乾补 feud.Barony = &乾补KannpurBarony{}

func init() {
    f := BKannpur乾补.(*乾补KannpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kannpur",
		TitleName: "乾补",
		TitleCode: "b_kannpur",
	}
}
