package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔西亚NursiaBarony struct {
	feud.BaseBarony
}

var BNursia努尔西亚 feud.Barony = &努尔西亚NursiaBarony{}

func init() {
	f := BNursia努尔西亚.(*努尔西亚NursiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nursia",
		TitleName: "努尔西亚",
		TitleCode: "b_nursia",
	}
}
