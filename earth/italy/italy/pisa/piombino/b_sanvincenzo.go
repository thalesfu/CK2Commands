package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣文琴佐SanvincenzoBarony struct {
	feud.BaseBarony
}

var BSanvincenzo圣文琴佐 feud.Barony = &圣文琴佐SanvincenzoBarony{}

func init() {
	f := BSanvincenzo圣文琴佐.(*圣文琴佐SanvincenzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanvincenzo",
		TitleName: "圣文琴佐",
		TitleCode: "b_sanvincenzo",
	}
}
