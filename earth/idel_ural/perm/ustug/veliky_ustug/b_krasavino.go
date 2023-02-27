package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉萨维诺KrasavinoBarony struct {
	feud.BaseBarony
}

var BKrasavino克拉萨维诺 feud.Barony = &克拉萨维诺KrasavinoBarony{}

func init() {
    f := BKrasavino克拉萨维诺.(*克拉萨维诺KrasavinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasavino",
		TitleName: "克拉萨维诺",
		TitleCode: "b_krasavino",
	}
}
