package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃默里希EmmerichBarony struct {
	feud.BaseBarony
}

var BEmmerich埃默里希 feud.Barony = &埃默里希EmmerichBarony{}

func init() {
	f := BEmmerich埃默里希.(*埃默里希EmmerichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emmerich",
		TitleName: "埃默里希",
		TitleCode: "b_emmerich",
	}
}
