package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 超戒寺VikramasilaBarony struct {
	feud.BaseBarony
}

var BVikramasila超戒寺 feud.Barony = &超戒寺VikramasilaBarony{}

func init() {
    f := BVikramasila超戒寺.(*超戒寺VikramasilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vikramasila",
		TitleName: "超戒寺",
		TitleCode: "b_vikramasila",
	}
}
