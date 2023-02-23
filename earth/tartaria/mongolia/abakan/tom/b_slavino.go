package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯拉维诺SlavinoBarony struct {
	feud.BaseBarony
}

var BSlavino斯拉维诺 feud.Barony = &斯拉维诺SlavinoBarony{}

func init() {
	f := BSlavino斯拉维诺.(*斯拉维诺SlavinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slavino",
		TitleName: "斯拉维诺",
		TitleCode: "b_slavino",
	}
}
