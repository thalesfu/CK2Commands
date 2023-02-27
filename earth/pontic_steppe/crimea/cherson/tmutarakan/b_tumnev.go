package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图姆涅夫TumnevBarony struct {
	feud.BaseBarony
}

var BTumnev图姆涅夫 feud.Barony = &图姆涅夫TumnevBarony{}

func init() {
    f := BTumnev图姆涅夫.(*图姆涅夫TumnevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tumnev",
		TitleName: "图姆涅夫",
		TitleCode: "b_tumnev",
	}
}
