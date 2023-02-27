package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多布良卡DobryankaBarony struct {
	feud.BaseBarony
}

var BDobryanka多布良卡 feud.Barony = &多布良卡DobryankaBarony{}

func init() {
    f := BDobryanka多布良卡.(*多布良卡DobryankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobryanka",
		TitleName: "多布良卡",
		TitleCode: "b_dobryanka",
	}
}
