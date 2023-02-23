package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣塞韦里诺SanseverinoBarony struct {
	feud.BaseBarony
}

var BSanseverino圣塞韦里诺 feud.Barony = &圣塞韦里诺SanseverinoBarony{}

func init() {
	f := BSanseverino圣塞韦里诺.(*圣塞韦里诺SanseverinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanseverino",
		TitleName: "圣塞韦里诺",
		TitleCode: "b_sanseverino",
	}
}
