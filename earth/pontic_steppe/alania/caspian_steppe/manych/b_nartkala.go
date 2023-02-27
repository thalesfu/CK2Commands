package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔特卡拉NartkalaBarony struct {
	feud.BaseBarony
}

var BNartkala纳尔特卡拉 feud.Barony = &纳尔特卡拉NartkalaBarony{}

func init() {
    f := BNartkala纳尔特卡拉.(*纳尔特卡拉NartkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nartkala",
		TitleName: "纳尔特卡拉",
		TitleCode: "b_nartkala",
	}
}
