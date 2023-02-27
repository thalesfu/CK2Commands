package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒苏Kyzyl_suuBarony struct {
	feud.BaseBarony
}

var BKyzyl_suu克孜勒苏 feud.Barony = &克孜勒苏Kyzyl_suuBarony{}

func init() {
    f := BKyzyl_suu克孜勒苏.(*克孜勒苏Kyzyl_suuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_suu",
		TitleName: "克孜勒苏",
		TitleCode: "b_kyzyl_suu",
	}
}
