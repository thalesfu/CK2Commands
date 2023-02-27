package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉凯TrakaiBarony struct {
	feud.BaseBarony
}

var BTrakai特拉凯 feud.Barony = &特拉凯TrakaiBarony{}

func init() {
    f := BTrakai特拉凯.(*特拉凯TrakaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trakai",
		TitleName: "特拉凯",
		TitleCode: "b_trakai",
	}
}
