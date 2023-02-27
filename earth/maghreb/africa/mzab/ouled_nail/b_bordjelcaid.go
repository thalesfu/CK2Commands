package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔吉采德BordjelcaidBarony struct {
	feud.BaseBarony
}

var BBordjelcaid布尔吉采德 feud.Barony = &布尔吉采德BordjelcaidBarony{}

func init() {
    f := BBordjelcaid布尔吉采德.(*布尔吉采德BordjelcaidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bordjelcaid",
		TitleName: "布尔吉采德",
		TitleCode: "b_bordjelcaid",
	}
}
