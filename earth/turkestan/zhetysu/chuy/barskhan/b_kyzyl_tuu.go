package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒图Kyzyl_tuuBarony struct {
	feud.BaseBarony
}

var BKyzyl_tuu克孜勒图 feud.Barony = &克孜勒图Kyzyl_tuuBarony{}

func init() {
    f := BKyzyl_tuu克孜勒图.(*克孜勒图Kyzyl_tuuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_tuu",
		TitleName: "克孜勒图",
		TitleCode: "b_kyzyl_tuu",
	}
}
