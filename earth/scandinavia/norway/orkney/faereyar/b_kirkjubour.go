package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇尔丘伯乌尔KirkjubourBarony struct {
	feud.BaseBarony
}

var BKirkjubour奇尔丘伯乌尔 feud.Barony = &奇尔丘伯乌尔KirkjubourBarony{}

func init() {
    f := BKirkjubour奇尔丘伯乌尔.(*奇尔丘伯乌尔KirkjubourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkjubour",
		TitleName: "奇尔丘伯乌尔",
		TitleCode: "b_kirkjubour",
	}
}
