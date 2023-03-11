package kyzyl-su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒苏Kyzyl-suBarony struct {
	feud.BaseBarony
}

var BKyzyl-su克孜勒苏 feud.Barony = &克孜勒苏Kyzyl-suBarony{}

func init() {
    f := BKyzyl-su克孜勒苏.(*克孜勒苏Kyzyl-suBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl-su",
		TitleName: "克孜勒苏",
		TitleCode: "b_kyzyl-su",
	}
}
