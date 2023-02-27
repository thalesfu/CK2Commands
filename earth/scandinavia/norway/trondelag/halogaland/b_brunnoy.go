package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布吕内于BrunnoyBarony struct {
	feud.BaseBarony
}

var BBrunnoy布吕内于 feud.Barony = &布吕内于BrunnoyBarony{}

func init() {
    f := BBrunnoy布吕内于.(*布吕内于BrunnoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brunnoy",
		TitleName: "布吕内于",
		TitleCode: "b_brunnoy",
	}
}
