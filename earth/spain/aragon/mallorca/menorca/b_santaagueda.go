package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿格达SantaaguedaBarony struct {
	feud.BaseBarony
}

var BSantaagueda圣阿格达 feud.Barony = &圣阿格达SantaaguedaBarony{}

func init() {
    f := BSantaagueda圣阿格达.(*圣阿格达SantaaguedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santaagueda",
		TitleName: "圣阿格达",
		TitleCode: "b_santaagueda",
	}
}
