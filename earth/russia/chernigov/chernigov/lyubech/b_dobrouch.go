package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多布鲁什DobrouchBarony struct {
	feud.BaseBarony
}

var BDobrouch多布鲁什 feud.Barony = &多布鲁什DobrouchBarony{}

func init() {
    f := BDobrouch多布鲁什.(*多布鲁什DobrouchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobrouch",
		TitleName: "多布鲁什",
		TitleCode: "b_dobrouch",
	}
}
