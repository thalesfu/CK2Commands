package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙姆希ShamshiBarony struct {
	feud.BaseBarony
}

var BShamshi沙姆希 feud.Barony = &沙姆希ShamshiBarony{}

func init() {
    f := BShamshi沙姆希.(*沙姆希ShamshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamshi",
		TitleName: "沙姆希",
		TitleCode: "b_shamshi",
	}
}
