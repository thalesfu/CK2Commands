package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅登布利克MedemblikBarony struct {
	feud.BaseBarony
}

var BMedemblik梅登布利克 feud.Barony = &梅登布利克MedemblikBarony{}

func init() {
    f := BMedemblik梅登布利克.(*梅登布利克MedemblikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medemblik",
		TitleName: "梅登布利克",
		TitleCode: "b_medemblik",
	}
}
