package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆哈拉SamhalBarony struct {
	feud.BaseBarony
}

var BSamhal萨姆哈拉 feud.Barony = &萨姆哈拉SamhalBarony{}

func init() {
	f := BSamhal萨姆哈拉.(*萨姆哈拉SamhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samhal",
		TitleName: "萨姆哈拉",
		TitleCode: "b_samhal",
	}
}
