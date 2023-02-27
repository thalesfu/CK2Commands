package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳伦卡拉NarinkalaBarony struct {
	feud.BaseBarony
}

var BNarinkala纳伦卡拉 feud.Barony = &纳伦卡拉NarinkalaBarony{}

func init() {
    f := BNarinkala纳伦卡拉.(*纳伦卡拉NarinkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narinkala",
		TitleName: "纳伦卡拉",
		TitleCode: "b_narinkala",
	}
}
