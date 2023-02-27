package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科里KoriBarony struct {
	feud.BaseBarony
}

var BKori科里 feud.Barony = &科里KoriBarony{}

func init() {
    f := BKori科里.(*科里KoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kori",
		TitleName: "科里",
		TitleCode: "b_kori",
	}
}
