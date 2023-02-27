package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔讷ElnaBarony struct {
	feud.BaseBarony
}

var BElna埃尔讷 feud.Barony = &埃尔讷ElnaBarony{}

func init() {
    f := BElna埃尔讷.(*埃尔讷ElnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elna",
		TitleName: "埃尔讷",
		TitleCode: "b_elna",
	}
}
