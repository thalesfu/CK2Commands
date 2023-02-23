package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔尼NarniBarony struct {
	feud.BaseBarony
}

var BNarni纳尔尼 feud.Barony = &纳尔尼NarniBarony{}

func init() {
	f := BNarni纳尔尼.(*纳尔尼NarniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narni",
		TitleName: "纳尔尼",
		TitleCode: "b_narni",
	}
}
