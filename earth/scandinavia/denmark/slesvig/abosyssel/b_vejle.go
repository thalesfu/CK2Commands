package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦埃勒VejleBarony struct {
	feud.BaseBarony
}

var BVejle瓦埃勒 feud.Barony = &瓦埃勒VejleBarony{}

func init() {
	f := BVejle瓦埃勒.(*瓦埃勒VejleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vejle",
		TitleName: "瓦埃勒",
		TitleCode: "b_vejle",
	}
}
