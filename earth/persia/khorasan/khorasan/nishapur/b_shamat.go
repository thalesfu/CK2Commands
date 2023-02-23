package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙马特ShamatBarony struct {
	feud.BaseBarony
}

var BShamat沙马特 feud.Barony = &沙马特ShamatBarony{}

func init() {
	f := BShamat沙马特.(*沙马特ShamatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamat",
		TitleName: "沙马特",
		TitleCode: "b_shamat",
	}
}
