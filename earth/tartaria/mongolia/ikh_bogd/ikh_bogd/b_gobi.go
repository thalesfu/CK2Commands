package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈壁GobiBarony struct {
	feud.BaseBarony
}

var BGobi戈壁 feud.Barony = &戈壁GobiBarony{}

func init() {
    f := BGobi戈壁.(*戈壁GobiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gobi",
		TitleName: "戈壁",
		TitleCode: "b_gobi",
	}
}
