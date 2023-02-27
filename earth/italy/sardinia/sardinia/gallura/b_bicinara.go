package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比西纳拉BicinaraBarony struct {
	feud.BaseBarony
}

var BBicinara比西纳拉 feud.Barony = &比西纳拉BicinaraBarony{}

func init() {
    f := BBicinara比西纳拉.(*比西纳拉BicinaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bicinara",
		TitleName: "比西纳拉",
		TitleCode: "b_bicinara",
	}
}
