package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑提耶SadiyaBarony struct {
	feud.BaseBarony
}

var BSadiya娑提耶 feud.Barony = &娑提耶SadiyaBarony{}

func init() {
    f := BSadiya娑提耶.(*娑提耶SadiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadiya",
		TitleName: "娑提耶",
		TitleCode: "b_sadiya",
	}
}
