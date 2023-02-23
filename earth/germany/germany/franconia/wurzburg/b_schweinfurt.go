package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施韦因富特SchweinfurtBarony struct {
	feud.BaseBarony
}

var BSchweinfurt施韦因富特 feud.Barony = &施韦因富特SchweinfurtBarony{}

func init() {
	f := BSchweinfurt施韦因富特.(*施韦因富特SchweinfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schweinfurt",
		TitleName: "施韦因富特",
		TitleCode: "b_schweinfurt",
	}
}
