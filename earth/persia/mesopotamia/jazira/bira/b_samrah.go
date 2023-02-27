package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨梅拉赫SamrahBarony struct {
	feud.BaseBarony
}

var BSamrah萨梅拉赫 feud.Barony = &萨梅拉赫SamrahBarony{}

func init() {
    f := BSamrah萨梅拉赫.(*萨梅拉赫SamrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samrah",
		TitleName: "萨梅拉赫",
		TitleCode: "b_samrah",
	}
}
