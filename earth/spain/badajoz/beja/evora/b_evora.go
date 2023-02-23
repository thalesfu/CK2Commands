package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃武拉EvoraBarony struct {
	feud.BaseBarony
}

var BEvora埃武拉 feud.Barony = &埃武拉EvoraBarony{}

func init() {
	f := BEvora埃武拉.(*埃武拉EvoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evora",
		TitleName: "埃武拉",
		TitleCode: "b_evora",
	}
}
