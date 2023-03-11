package kyzyl_su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒苏Kyzyl_suBarony struct {
	feud.BaseBarony
}

var BKyzyl_su克孜勒苏 feud.Barony = &克孜勒苏Kyzyl_suBarony{}

func init() {
    f := BKyzyl_su克孜勒苏.(*克孜勒苏Kyzyl_suBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_su",
		TitleName: "克孜勒苏",
		TitleCode: "b_kyzyl-su",
	}
}
