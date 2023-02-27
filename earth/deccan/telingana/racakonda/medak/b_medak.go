package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥陀迦MedakBarony struct {
	feud.BaseBarony
}

var BMedak弥陀迦 feud.Barony = &弥陀迦MedakBarony{}

func init() {
    f := BMedak弥陀迦.(*弥陀迦MedakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medak",
		TitleName: "弥陀迦",
		TitleCode: "b_medak",
	}
}
