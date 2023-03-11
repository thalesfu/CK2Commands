package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃罗诺沃VoronovoBarony struct {
	feud.BaseBarony
}

var BVoronovo沃罗诺沃 feud.Barony = &沃罗诺沃VoronovoBarony{}

func init() {
    f := BVoronovo沃罗诺沃.(*沃罗诺沃VoronovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voronovo",
		TitleName: "沃罗诺沃",
		TitleCode: "b_voronovo",
	}
}
