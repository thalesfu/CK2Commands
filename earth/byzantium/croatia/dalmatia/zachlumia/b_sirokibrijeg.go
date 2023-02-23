package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希罗基布里耶格SirokibrijegBarony struct {
	feud.BaseBarony
}

var BSirokibrijeg希罗基布里耶格 feud.Barony = &希罗基布里耶格SirokibrijegBarony{}

func init() {
	f := BSirokibrijeg希罗基布里耶格.(*希罗基布里耶格SirokibrijegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirokibrijeg",
		TitleName: "希罗基布里耶格",
		TitleCode: "b_sirokibrijeg",
	}
}
