package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣布里厄St_brieucBarony struct {
	feud.BaseBarony
}

var BSt_brieuc圣布里厄 feud.Barony = &圣布里厄St_brieucBarony{}

func init() {
    f := BSt_brieuc圣布里厄.(*圣布里厄St_brieucBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_brieuc",
		TitleName: "圣布里厄",
		TitleCode: "b_st_brieuc",
	}
}
