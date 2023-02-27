package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伊马济TuymazyBarony struct {
	feud.BaseBarony
}

var BTuymazy图伊马济 feud.Barony = &图伊马济TuymazyBarony{}

func init() {
    f := BTuymazy图伊马济.(*图伊马济TuymazyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuymazy",
		TitleName: "图伊马济",
		TitleCode: "b_tuymazy",
	}
}
