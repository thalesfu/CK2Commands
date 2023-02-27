package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯利恒BethlehemBarony struct {
	feud.BaseBarony
}

var BBethlehem伯利恒 feud.Barony = &伯利恒BethlehemBarony{}

func init() {
    f := BBethlehem伯利恒.(*伯利恒BethlehemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bethlehem",
		TitleName: "伯利恒",
		TitleCode: "b_bethlehem",
	}
}
