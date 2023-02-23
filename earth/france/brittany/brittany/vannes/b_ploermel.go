package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普洛埃梅勒PloermelBarony struct {
	feud.BaseBarony
}

var BPloermel普洛埃梅勒 feud.Barony = &普洛埃梅勒PloermelBarony{}

func init() {
	f := BPloermel普洛埃梅勒.(*普洛埃梅勒PloermelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ploermel",
		TitleName: "普洛埃梅勒",
		TitleCode: "b_ploermel",
	}
}
