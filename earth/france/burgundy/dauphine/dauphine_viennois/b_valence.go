package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦朗斯ValenceBarony struct {
	feud.BaseBarony
}

var BValence瓦朗斯 feud.Barony = &瓦朗斯ValenceBarony{}

func init() {
    f := BValence瓦朗斯.(*瓦朗斯ValenceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valence",
		TitleName: "瓦朗斯",
		TitleCode: "b_valence",
	}
}
