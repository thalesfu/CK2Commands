package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶灵JellingBarony struct {
	feud.BaseBarony
}

var BJelling耶灵 feud.Barony = &耶灵JellingBarony{}

func init() {
	f := BJelling耶灵.(*耶灵JellingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jelling",
		TitleName: "耶灵",
		TitleCode: "b_jelling",
	}
}
